package main

import (
	"fmt"
	// "time" // No longer needed

	cr "mediator_pattern_chat_room_go/chat_room"
)

func main() {
	// Match Python's initial print
	fmt.Println("--- Mediator Pattern: Chat Room Demo ---")

	// 1. Create the Mediator (Chat Room)
	mediator := cr.NewChatRoom()

	// 2. Create Colleagues (Users)
	alice := cr.NewChatUser("Alice")
	bob := cr.NewChatUser("Bob")
	charlie := cr.NewChatUser("Charlie")

	// 3. Register users with the Mediator (Match Python structure)
	fmt.Println("\n--- Adding users to the chat room ---")
	mediator.AddUser(alice)
	mediator.AddUser(bob)
	mediator.AddUser(charlie)
	// Try adding an existing user (like Python)
	mediator.AddUser(bob) // This will now print the "already in room" message

	// 4. Users send messages through the Mediator (Match Python messages and structure)
	fmt.Println("\n--- Users sending messages ---")
	alice.Send("Hi everyone! How's it going?")
	fmt.Println("--------------------") // Match Python separator
	bob.Send("Hey Alice! Doing well, thanks. Just working on the project.")
	fmt.Println("--------------------") // Match Python separator
	charlie.Send("Hello! Project is coming along nicely.")

	// 5. Remove a user (Match Python structure)
	fmt.Println("\n--- Removing a user ---")
	mediator.RemoveUser(bob)

	// 6. Send another message (Match Python message)
	fmt.Println("\n--- Sending message after Bob left ---")
	alice.Send("Okay, let's sync up later then.")

	// 7. Try removing a non-existent user (Match Python structure)
	fmt.Println("\n--- Trying to remove Bob again ---")
	mediator.RemoveUser(bob) // Attempt to remove Bob again, will print "not in room"

	// Match Python's final print
	fmt.Println("\n--- Chat Room Demo Complete ---")
}
