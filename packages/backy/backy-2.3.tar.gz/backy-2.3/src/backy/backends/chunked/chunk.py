import binascii
import io
import lzo
import mmh3
import os
import tempfile
import time


chunk_stats = {
    'write_full': 0,
    'write_partial': 0,
}


class Chunk(object):
    """A chunk in a file that represents a part of it.

    This can be read and updated in memory, which updates its
    hash and then can be flushed to disk when appropriate.

    """

    CHUNK_SIZE = 4 * 1024**2  # 4 MiB chunks

    _read_existing_called = False  # Test support

    def __init__(self, file, id, store, hash):
        self.id = id
        self.hash = hash
        self.file = file
        self.store = store
        self.clean = True
        self.loaded = False

        if self.id not in file._access_stats:
            self.file._access_stats[id] = (0, 0)

        self.data = None

    def _cache_prio(self):
        return self.file._access_stats[self.id]

    def _touch(self):
        count = self.file._access_stats[self.id][0]
        self.file._access_stats[self.id] = (count + 1, time.time())

    def _read_existing(self):
        self._read_existing_called = True  # Test support
        if self.data is not None:
            return
        # Prepare working with the chunk. We keep the data in RAM for
        # easier random access combined with transparent compression.
        data = b''
        if self.hash:
            chunk_file = self.store.chunk_path(self.hash)
            with open(chunk_file, 'rb') as f:
                data = f.read()
                data = lzo.decompress(data)
            disk_hash = hash(data)
            # This is a safety belt. Hashing is sufficiently fast to avoid
            # us accidentally reading garbage.
            if disk_hash != self.hash:
                raise ValueError("Expected hash {} but data has {}".format(
                    self.hash, disk_hash))
        self._init_data(data)

    def _init_data(self, data):
        self.data = io.BytesIO(data)

    def read(self, offset, size=-1):
        """Read data from the chunk.

        Return the data and the remaining size that should be read.
        """
        self._read_existing()
        self._touch()

        self.data.seek(offset)
        data = self.data.read(size)
        remaining = -1
        if size != -1:
            remaining = max([0, size - len(data)])
        return data, remaining

    def write(self, offset, data):
        """Write data to the chunk, returns

        - the amount of data we used
        - the _data_ remaining

        """
        self._touch()

        remaining_data = data[self.CHUNK_SIZE - offset:]
        data = data[:self.CHUNK_SIZE - offset]

        if offset == 0 and len(data) == self.CHUNK_SIZE:
            # Special case: overwrite the entire chunk.
            self._init_data(data)
            chunk_stats['write_full'] += 1
        else:
            self._read_existing()
            self.data.seek(offset)
            self.data.write(data)
            chunk_stats['write_partial'] += 1
        self.clean = False

        return len(data), remaining_data

    def flush(self):
        if self.clean:
            return
        assert self.data is not None
        self._update_hash()
        target = self.store.chunk_path(self.hash)
        needs_forced_write = (
            self.store.force_writes and
            self.hash not in self.store.seen_forced)
        if not os.path.exists(target) or needs_forced_write:
            fd, tmpfile_name = tempfile.mkstemp(dir=self.store.path)
            with os.fdopen(fd, mode='wb') as f:
                data = lzo.compress(self.data.getvalue())
                f.write(data)
            subdir = os.path.dirname(target)
            try:
                os.makedirs(subdir)
            except FileExistsError:
                # Someone else accessing the store may have created this
                # already.
                pass
            os.rename(tmpfile_name, target)
            os.chmod(target, 0o440)
            self.store.seen_forced.add(self.hash)
        self.clean = True

    def _update_hash(self):
        # I'm not using read() here to a) avoid cache accounting and b)
        # use a faster path to get the data.
        data = self.data.getvalue()
        self.hash = hash(data)
        self.file._mapping[self.id] = self.hash


def hash(data):
    return binascii.hexlify(mmh3.hash_bytes(data)).decode('ascii')
