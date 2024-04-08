import EventHandlerInterface from "../../../@shared/event/event-handler.interface";
import CustomerChangedAddresEvent from "../customer-changed-address.event";

export default class EnviaConsolelogHandler implements 
    EventHandlerInterface<CustomerChangedAddresEvent>
{
    handle(event: CustomerChangedAddresEvent): void {
       console.log(`Endereco do cliente: ${event.eventData.id},
       ${event.eventData.name} alterado para: ${event.eventData.endereco}`)
    }
}