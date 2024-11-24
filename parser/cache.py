import os
import json

from config import CACHE_PATH

CACHE_FILE = os.path.join(CACHE_PATH, "cache.json")
os.makedirs(CACHE_PATH, exist_ok=True)


class Cache:
    def __init__(self):
        self.file_path = CACHE_FILE
        self.data = self._load_cache()

    def _load_cache(self):
        if os.path.exists(self.file_path):
            with open(self.file_path, "r", encoding="utf-8") as f:
                return json.load(f)
        return {}

    def save(self):
        with open(self.file_path, "w", encoding="utf-8") as f:
            json.dump(self.data, f)

    def get(self, key, default=None):
        return self.data.get(key, default)

    def set(self, key, value):
        self.data[key] = value
        self.save()

    def delete(self, key):
        if key in self.data:
            del self.data[key]
            self.save()


cache = Cache()
