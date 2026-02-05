import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  async getHtml() {
    return /*html*/ `
      <section class="home-view">
        <dialog id="assign-tenant-to-slot-modal" closedby="any">
          <header>
            <h3>Reservar plaza</h3>
            <close-modal-button></close-modal-button>
          </header>
          <slot-form></slot-form>
        </dialog>
  
        <header class="section-header">
          <h2>¡Bienvenido!</h2>
        </header>
        
        <slots-collection></slots-collection>
        <div></div>
      </section>
    `;
  }

  setUpJavascript() {
    const view = document.querySelector(".home-view");
    const modal = document.querySelector("#assign-tenant-to-slot-modal");
    const form = document.querySelector("slot-form");

    view.addEventListener("slot:selected", (event) => {
      form.slotId = event.detail.slotId;
      form.slotNumber = event.detail.slotNumber;
      modal.showModal();
    });

    const closeBehavior = () => {
      modal.close();
      form.clear();
    };

    view.addEventListener("close-modal", closeBehavior);
    view.addEventListener("slot:assigned", closeBehavior);
  }
}
