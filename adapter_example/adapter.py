from typing import Final

from target import Target
from adaptee import Adaptee


class Adapter(Target):
    def __init__(self, adaptee: Adaptee) -> None:
        self._adaptee: Final[Adaptee] = adaptee

    def request(self) -> str:
        translated = self._adaptee.specific_request()[::-1]
        return f"Adapter: (TRANSLATED) {translated}"