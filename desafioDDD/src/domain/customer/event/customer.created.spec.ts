import EventDispatcher from "../../@shared/event/event-dispatcher"
import CustomerCreatedEvent from "./customer-created.event";
import EnviaConsolelog1Handler from "./handler/envia-console-log1.handler";
import EnviaConsolelog2Handler from "./handler/envia-console-log2.handler";

describe("Customer created event Unit Tests", () => {

    it("Should run console.log1 when Customer is created.",() => {
        const eventDispatcher = new EventDispatcher();
        const eventHandler = new EnviaConsolelog1Handler();
    
        eventDispatcher.register("CustomerCreatedEvent", eventHandler);
    
        expect(eventDispatcher.getEventHandlers
            ["CustomerCreatedEvent"]).toBeDefined();
    
        expect(eventDispatcher.getEventHandlers
            ["CustomerCreatedEvent"].length).toBe(1);

        expect(eventDispatcher.getEventHandlers
            ["CustomerCreatedEvent"][0]).toMatchObject(eventHandler);
            
        const spyEventHandler = jest.spyOn(eventHandler, "handle");
        const customerCreatedEvent = new CustomerCreatedEvent({
            changed: "Name",
            customer_id: "abc",
        });

        eventDispatcher.notify(customerCreatedEvent);
        expect(spyEventHandler).toBeCalledTimes(1);

    });

    it("Should run console.log2 when Customer is created.",() => {
        const eventDispatcher = new EventDispatcher();
        const eventHandler = new EnviaConsolelog2Handler();
    
        eventDispatcher.register("CustomerCreatedEvent", eventHandler);
    
        expect(eventDispatcher.getEventHandlers
            ["CustomerCreatedEvent"]).toBeDefined();
    
        expect(eventDispatcher.getEventHandlers
            ["CustomerCreatedEvent"].length).toBe(1);

        expect(eventDispatcher.getEventHandlers
            ["CustomerCreatedEvent"][0]).toMatchObject(eventHandler);
            
        const spyEventHandler = jest.spyOn(eventHandler, "handle");
        const customerCreatedEvent = new CustomerCreatedEvent({
            changed: "Name",
            customer_id: "abc",
        });

        eventDispatcher.notify(customerCreatedEvent);
        expect(spyEventHandler).toBeCalledTimes(1);
    });

    it("Should run console.log1 and console.log2 when Customer is created.",() => {
        const eventDispatcher = new EventDispatcher();
        const eventHandler1 = new EnviaConsolelog1Handler();
        const eventHandler2 = new EnviaConsolelog2Handler(); 

        eventDispatcher.register("CustomerCreatedEvent", eventHandler1);
        eventDispatcher.register("CustomerCreatedEvent", eventHandler2);
    
        expect(eventDispatcher.getEventHandlers
            ["CustomerCreatedEvent"]).toBeDefined();
    
        expect(eventDispatcher.getEventHandlers
            ["CustomerCreatedEvent"].length).toBe(2);

        expect(eventDispatcher.getEventHandlers
            ["CustomerCreatedEvent"][0]).toMatchObject(eventHandler1);
        expect(eventDispatcher.getEventHandlers
            ["CustomerCreatedEvent"][1]).toMatchObject(eventHandler2);
            
        const spyEventHandler1 = jest.spyOn(eventHandler1, "handle");
        const spyEventHandler2 = jest.spyOn(eventHandler2, "handle");
        const customerCreatedEvent = new CustomerCreatedEvent({
            changed: "Name",
            customer_id: "abc",
        });

        eventDispatcher.notify(customerCreatedEvent);
        expect(spyEventHandler1).toBeCalledTimes(1);
        expect(spyEventHandler2).toBeCalledTimes(1);
    });


    
})