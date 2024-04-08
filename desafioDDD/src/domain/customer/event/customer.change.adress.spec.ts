import EventDispatcher from "../../@shared/event/event-dispatcher"
import CustomerChangedAddresEvent from "./customer-changed-address.event";
import EnviaConsolelogHandler from "./handler/envia-console-log.handler";

describe("Customer changed Address event Unit tests ", () => {

    it("Should run console.log when address is changed.", () => {
        const eventDispatcher = new EventDispatcher();
        const eventHandler = new EnviaConsolelogHandler();

        const changedAddressEvent = new CustomerChangedAddresEvent(
            {
             id: "MockedCustomerid",
             name: "MockedCustomerName",
             endereco: "Novo endereco",
            }
        );

        eventDispatcher.register("CustomerChangedAddresEvent", eventHandler);
        expect(eventDispatcher.getEventHandlers
            ["CustomerChangedAddresEvent"]).toBeDefined();
    
        expect(eventDispatcher.getEventHandlers
            ["CustomerChangedAddresEvent"].length).toBe(1);

        expect(eventDispatcher.getEventHandlers
            ["CustomerChangedAddresEvent"][0]).toMatchObject(eventHandler);
        
        const spy = jest.spyOn(eventHandler, "handle");
        eventDispatcher.notify(changedAddressEvent);
        expect(spy).toBeCalledTimes(1);
    })
})