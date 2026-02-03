import SlotCard from "./SlotCard.js";

const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    div {
      display: grid;
      grid-template-columns: repeat(4, minmax(10rem, 1fr));
      gap: 1rem;
      place-items: center;
    }
  </style>

  <div></div>
`;

export default class SlotsCollection extends HTMLElement {
  tenants = [];

  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
    this.slots = [];
  }

  async handleEvent() {
    this.renderPlaceholders();
    await this.fetchSlots();
    this.render();
  }

  async connectedCallback() {
    if (this.slots.length != 0) return;

    document.addEventListener("slot:assigned", this);

    this.renderPlaceholders();
    await this.fetchSlots();
    this.render();
  }

  async fetchSlots() {
    const SLOTS_URL = import.meta.env.VITE_API_URL + "/api/slots";

    await fetch(SLOTS_URL)
      .then((response) => response.json())
      .then((json) => {
        this.slots = json["data"];
      })
      .catch((error) => {
        console.error("Error fetching slots:", error);
      });
  }

  renderPlaceholders() {
    const container = this.shadowRoot.querySelector("div");
    container.innerHTML = "";
    const COLLECTION_LENGTH = 12;

    for (let i = 0; i < COLLECTION_LENGTH; i++) {
      container.appendChild(new SlotCard());
    }
  }

  render() {
    const container = this.shadowRoot.querySelector("div");
    container.innerHTML = "";

    this.slots.forEach((slot) => {
      const card = SlotCard.fromJson(slot);
      container.appendChild(card);
    });
  }
}

customElements.define("slots-collection", SlotsCollection);
