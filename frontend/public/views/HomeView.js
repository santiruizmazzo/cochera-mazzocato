import AbstractView from "./AbstractView.js";
import ContentSection from "../components/ContentSection.js";
import SlotsCollection from "../components/SlotsCollection.js";
import SlotForm from "../components/SlotForm.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  renderWithin(mainElement) {
    this.fetchSlots();

    this.homeSection = new ContentSection();

    this.homeSection.title = "¡Bienvenido a la Cochera Mazzocato!";
    this.homeSection.hideButton = true;

    this.slotForm = new SlotForm();
    this.homeSection.modalTitle = "Reservar plaza";
    this.homeSection.modalContent = this.slotForm;

    this.slotsCollection = new SlotsCollection();
    this.homeSection.content = this.slotsCollection;

    mainElement.appendChild(this.homeSection);

    this.setUpEventListeners();
  }

  async fetchSlots() {
    const SLOTS_URL = import.meta.env.VITE_API_URL + "/api/slots";

    await fetch(SLOTS_URL)
      .then((response) => {
        if (!response.ok) {
          throw new Error(`Error ${response.status}`);
        }
        return response.json();
      })
      .then(({ data }) => {
        this.slotsCollection.load(data);
      })
      .catch((error) => {
        this.homeSection.errorMessage = error;
      });
  }

  setUpEventListeners() {
    this.homeSection.addEventListener("slot:assigned", () => {
      this.homeSection.modal.close();
      this.fetchSlots();
    });
  }
}
