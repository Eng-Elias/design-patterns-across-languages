// chat_room.test.ts

import { ChatRoom, ChatUser, User, ChatMediator } from './chat_room';

// Mocking the User interface for isolated Mediator testing
const mockUser = (name: string): User => ({
    getName: jest.fn().mockReturnValue(name),
    send: jest.fn(),
    receive: jest.fn(),
    setMediator: jest.fn(), // Important to mock this method used in addUser
});

describe('ChatRoom Mediator', () => {
    let mediator: ChatRoom;
    let mockAlice: User;
    let mockBob: User;

    beforeEach(() => {
        // Reset console mocks before each test if needed
        // jest.spyOn(console, 'log').mockImplementation(() => {}); // Suppress logs if desired

        mediator = new ChatRoom();
        mockAlice = mockUser("Alice");
        mockBob = mockUser("Bob");
    });

    afterEach(() => {
       // jest.restoreAllMocks(); // Restore original console.log etc.
    });

    test('should add users correctly and set their mediator', () => {
        mediator.addUser(mockAlice);
        mediator.addUser(mockBob);

        // Check internal state (usually discouraged, but okay for demonstration/simpler tests)
        // @ts-ignore Accessing private member for test verification
        expect(mediator.users.size).toBe(2);
        // @ts-ignore
        expect(mediator.users.get("Alice")).toBe(mockAlice);
        // @ts-ignore
        expect(mediator.users.get("Bob")).toBe(mockBob);

        // Verify that the mediator was set for each user
        expect(mockAlice.setMediator).toHaveBeenCalledWith(mediator);
        expect(mockBob.setMediator).toHaveBeenCalledWith(mediator);
    });

    test('should not add the same user twice', () => {
        mediator.addUser(mockAlice);
        // @ts-ignore
        expect(mediator.users.size).toBe(1);

        mediator.addUser(mockAlice); // Add again
        // @ts-ignore
        expect(mediator.users.size).toBe(1); // Size should remain 1
        expect(mockAlice.setMediator).toHaveBeenCalledTimes(1); // setMediator only called once
    });

    test('should remove users correctly', () => {
        mediator.addUser(mockAlice);
        // @ts-ignore
        expect(mediator.users.has("Alice")).toBe(true);

        mediator.removeUser(mockAlice);
        // @ts-ignore
        expect(mediator.users.has("Alice")).toBe(false);
        // @ts-ignore
        expect(mediator.users.size).toBe(0);
    });

    test('should not fail when removing a non-existent user', () => {
        const mockCharlie = mockUser("Charlie");
        expect(() => mediator.removeUser(mockCharlie)).not.toThrow();
         // @ts-ignore
        expect(mediator.users.size).toBe(0);
    });

    test('should send message to all users except the sender', () => {
        const userAlice = new ChatUser("Alice"); // Use real users for integration test
        const userBob = new ChatUser("Bob");
        const userCharlie = new ChatUser("Charlie");

        // Spy on the receive method of real users
        const receiveSpyAlice = jest.spyOn(userAlice, 'receive');
        const receiveSpyBob = jest.spyOn(userBob, 'receive');
        const receiveSpyCharlie = jest.spyOn(userCharlie, 'receive');

        mediator.addUser(userAlice);
        mediator.addUser(userBob);
        mediator.addUser(userCharlie);

        const messageAlice = "Hello from Alice!";
        userAlice.send(messageAlice); // Alice sends the message

        // Verify receive calls
        expect(receiveSpyAlice).not.toHaveBeenCalled();
        expect(receiveSpyBob).toHaveBeenCalledWith(messageAlice, "Alice");
        expect(receiveSpyCharlie).toHaveBeenCalledWith(messageAlice, "Alice");
        expect(receiveSpyBob).toHaveBeenCalledTimes(1);
        expect(receiveSpyCharlie).toHaveBeenCalledTimes(1);

        // Clear mocks before the next send
        receiveSpyAlice.mockClear();
        receiveSpyBob.mockClear();
        receiveSpyCharlie.mockClear();

        // Bob sends a message
        const messageBob = "Hi back!";
        userBob.send(messageBob);

        expect(receiveSpyAlice).toHaveBeenCalledWith(messageBob, "Bob");
        expect(receiveSpyBob).not.toHaveBeenCalled();
        expect(receiveSpyCharlie).toHaveBeenCalledWith(messageBob, "Bob");
        expect(receiveSpyAlice).toHaveBeenCalledTimes(1);
        expect(receiveSpyCharlie).toHaveBeenCalledTimes(1);
    });
});

describe('ChatUser Colleague', () => {
     test('send method should call mediator.sendMessage', () => {
        // Mock the mediator
        const mockMediator: ChatMediator = {
            sendMessage: jest.fn(),
            addUser: jest.fn(),
            removeUser: jest.fn(),
        };

        const userDave = new ChatUser("Dave");
        userDave.setMediator(mockMediator); // Set the mock mediator

        const message = "Test message";
        userDave.send(message);

        // Verify the mediator's method was called correctly
        expect(mockMediator.sendMessage).toHaveBeenCalledWith(message, userDave);
        expect(mockMediator.sendMessage).toHaveBeenCalledTimes(1);
    });

     test('send method should handle missing mediator', () => {
        const userEve = new ChatUser("Eve");
        // Mediator is not set
         const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {}); // Suppress error log

        expect(() => userEve.send("Message without mediator")).not.toThrow();
        expect(errorSpy).toHaveBeenCalledWith("'Eve' cannot send message: Mediator not set.");

        errorSpy.mockRestore(); // Clean up spy
    });
});