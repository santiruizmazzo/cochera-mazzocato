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
    this.shadowRoot
      .querySelector("div")
      .addEventListener("slot:assigned", this);
    this.slots = [];
  }

  handleEvent(event) {
    this.renderPlaceholders();
    this.dispatchEvent(event);
  }

  connectedCallback() {
    if (this.slots.length === 0) {
      this.renderPlaceholders();
      return;
    }
    this.render();
  }

  renderPlaceholders() {
    const container = this.shadowRoot.querySelector("div");
    container.innerHTML = "";
    const COLLECTION_LENGTH = 12;

    for (let i = 0; i < COLLECTION_LENGTH; i++) {
      container.appendChild(new SlotCard());
    }
  }

  loadSlots(slots) {
    this.slots = slots;
    this.render();
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
