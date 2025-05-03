import { ChatRoom, ChatUser } from './chat_room';

function main() {
    console.log("--- Mediator Pattern: Chat Room Demo (TypeScript) ---");

    // 1. Create the Mediator (Chat Room)
    const mediator = new ChatRoom();

    // 2. Create Colleagues (Users)
    // Note: Mediator is set when adding the user to the room in this implementation
    const userAlice = new ChatUser("Alice");
    const userBob = new ChatUser("Bob");
    const userCharlie = new ChatUser("Charlie");

    // 3. Register users with the Mediator
    console.log("\n--- Adding users to the chat room ---");
    mediator.addUser(userAlice);
    mediator.addUser(userBob);
    mediator.addUser(userCharlie);
    // Try adding an existing user
    mediator.addUser(userBob); 

    // 4. Users send messages through the Mediator
    console.log("\n--- Users sending messages ---");
    userAlice.send("Hi everyone! How's it going?");
    console.log("-".repeat(20));
    userBob.send("Hey Alice! Doing well, thanks. Just working on the project.");
    console.log("-".repeat(20));
    userCharlie.send("Hello! Project is coming along nicely.");

    // 5. Remove a user
    console.log("\n--- Removing a user ---");
    mediator.removeUser(userBob);

    // 6. Send another message
    console.log("\n--- Sending message after Bob left ---");
    userAlice.send("Okay, let's sync up later then.");

    // 7. Try removing a non-existent user
    console.log("\n--- Trying to remove Bob again ---");
    mediator.removeUser(userBob);

    console.log("\n--- Chat Room Demo Complete ---");
}

// Run the demo
main();