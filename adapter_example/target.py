from abc import ABC, abstractmethod


class Target(ABC):
    @abstractmethod
    def request(self) -> str:
        """Return a string in the format the client expects."""
        raise NotImplementedError