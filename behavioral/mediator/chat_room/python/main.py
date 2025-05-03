from chat_room import ChatRoom, ChatUser

def main():
    """
    Client code: Demonstrates the use of the Mediator pattern
    for a chat room application.
    """
    print("--- Mediator Pattern: Chat Room Demo ---")

    # 1. Create the Mediator (Chat Room)
    mediator = ChatRoom()

    # 2. Create Colleagues (Users)
    user_alice = ChatUser(mediator, "Alice")
    user_bob = ChatUser(mediator, "Bob")
    user_charlie = ChatUser(mediator, "Charlie")

    # 3. Register users with the Mediator
    print("\n--- Adding users to the chat room ---")
    mediator.add_user(user_alice)
    mediator.add_user(user_bob)
    mediator.add_user(user_charlie)
    # Try adding an existing user
    mediator.add_user(user_bob) 

    # 4. Users send messages through the Mediator
    print("\n--- Users sending messages ---")
    user_alice.send("Hi everyone! How's it going?")
    print("-" * 20)
    user_bob.send("Hey Alice! Doing well, thanks. Just working on the project.")
    print("-" * 20)
    user_charlie.send("Hello! Project is coming along nicely.")

    # 5. Remove a user
    print("\n--- Removing a user ---")
    mediator.remove_user(user_bob)

    # 6. Send another message
    print("\n--- Sending message after Bob left ---")
    user_alice.send("Okay, let's sync up later then.")

    # 7. Try removing a non-existent user
    print("\n--- Trying to remove Bob again ---")
    mediator.remove_user(user_bob)

    print("\n--- Chat Room Demo Complete ---")

if __name__ == "__main__":
    main()