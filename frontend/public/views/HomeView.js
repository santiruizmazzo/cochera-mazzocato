import AbstractView from "./AbstractView.js";
import ContentSection from "../components/ContentSection.js";
import SlotsCollection from "../components/SlotsCollection.js";
import SlotForm from "../components/SlotForm.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  renderWithin(mainElement) {
    this.fetchSlots().then((slots) => {
      this.slotsCollection.loadSlots(slots);
    });

    this.homeSection = new ContentSection();

    this.homeSection.title = "¡Bienvenido a la Cochera Mazzocato!";
    this.homeSection.hideButton = true;

    this.slotForm = new SlotForm();
    this.homeSection.modalContent = this.slotForm;

    this.slotsCollection = new SlotsCollection();
    this.homeSection.content = this.slotsCollection;

    mainElement.appendChild(this.homeSection);

    this.setUpJavascript();
  }

  async fetchSlots() {
    const SLOTS_URL = import.meta.env.VITE_API_URL + "/api/slots";

    return await fetch(SLOTS_URL)
      .then((response) => response.json())
      .then((json) => {
        return json["data"];
      })
      .catch((error) => {
        console.error("Error fetching slots:", error);
      });
  }

  async getHtml() {
    return /*html*/ `
      <section class="home-view">
        <custom-modal title="Reservar plaza">
          <slot-form></slot-form>
        </custom-modal>
  
        <header class="section-header">
          <h2>¡Bienvenido!</h2>
        </header>
        
        <slots-collection></slots-collection>
        <div></div>
      </section>
    `;
  }

  setUpJavascript() {
    this.homeSection.addEventListener("slot:assigned", () => {
      this.homeSection.modal.close();
      this.fetchSlots().then((slots) => {
        this.slotsCollection.loadSlots(slots);
      });
    });
  }
}
