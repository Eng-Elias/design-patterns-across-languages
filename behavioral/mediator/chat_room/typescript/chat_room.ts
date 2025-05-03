// chat_room.ts

// --- Colleague Interface ---
// Forward declaration for Mediator is tricky in TS without complex setups,
// so we'll define User first and ensure Mediator is defined before use.

/**
 * Represents a user in the chat system.
 * Users communicate only through a Mediator.
 */
export interface User {
    getName(): string;
    send(message: string): void;
    receive(message: string, senderName: string): void;
    // Method to set the mediator after initialization if needed, or pass in constructor
    setMediator(mediator: ChatMediator): void; 
}


// --- Mediator Interface ---

/**
 * The Mediator interface declares methods used by Users
 * to notify the mediator about events.
 */
export interface ChatMediator {
    sendMessage(message: string, sender: User): void;
    addUser(user: User): void;
    removeUser(user: User): void;
}


// --- Concrete Colleague ---

/**
 * Concrete Colleague (User) implementation.
 */
export class ChatUser implements User {
    private mediator: ChatMediator | null = null; // Mediator can be set later
    private name: string;

    constructor(name: string) {
        this.name = name;
        // Mediator might be set after creation depending on setup
    }

    // Allows setting the mediator after instantiation
    public setMediator(mediator: ChatMediator): void {
        this.mediator = mediator;
    }

    public getName(): string {
        return this.name;
    }

    public send(message: string): void {
        if (!this.mediator) {
            console.error(`'${this.name}' cannot send message: Mediator not set.`);
            return;
        }
        console.log(`'${this.name}' preparing to send message: '${message}'`);
        this.mediator.sendMessage(message, this);
    }

    public receive(message: string, senderName: string): void {
        console.log(`'${this.name}' received message from '${senderName}': '${message}'`);
    }
}


// --- Concrete Mediator ---

/**
 * Concrete Mediator implementation for the chat room.
 * Coordinates communication between ChatUsers.
 */
export class ChatRoom implements ChatMediator {
    // Using Map for easier user lookup and removal by name
    private users: Map<string, User> = new Map();

    constructor() {
        console.log("ChatRoom created.");
    }

    public addUser(user: User): void {
        const userName = user.getName();
        if (!this.users.has(userName)) {
            this.users.set(userName, user);
            user.setMediator(this); // Set the mediator for the user
            console.log(`'${userName}' joined the chat room.`);
        } else {
            console.log(`User '${userName}' is already in the chat room.`);
        }
    }

    public removeUser(user: User): void {
        const userName = user.getName();
        if (this.users.has(userName)) {
            this.users.delete(userName);
             // Optionally, nullify the mediator reference in the user
             // user.setMediator(null); // Depends on desired behavior
            console.log(`'${userName}' left the chat room.`);
        } else {
            console.log(`User '${userName}' is not in the chat room.`);
        }
    }

    public sendMessage(message: string, sender: User): void {
        const senderName = sender.getName();
        console.log(`'${senderName}' sends message: '${message}'`);
        this.users.forEach((user, name) => {
            // Send message to everyone except the sender
            if (user !== sender) {
                user.receive(message, senderName);
            }
        });
    }
}