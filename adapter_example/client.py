from target import Target
from adapter import Adapter
from adaptee import Adaptee


def client_code(target: Target) -> None:
    print(target.request())


if __name__ == "__main__":
    print("Client: The Adaptee has a weird interface:")
    adaptee = Adaptee()
    print(f"Adaptee: {adaptee.specific_request()}")

    print("\nClient: But with an adapter, I can work with it:")
    adapter = Adapter(adaptee)
    client_code(adapter)