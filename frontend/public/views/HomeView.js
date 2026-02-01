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
            <button class="close-action-btn close-modal-btn">
              <svg><use href="#close"/></svg>
            </button>
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
    const modal = document.querySelector("#assign-tenant-to-slot-modal");
    const closeModal = document.querySelector(".close-modal-btn");
    const form = document.querySelector("slot-form");
    const view = document.querySelector(".home-view");

    view.addEventListener("slot:selected", (event) => {
      form.slotId = event.detail.slotId;
      form.slotNumber = event.detail.slotNumber;
      modal.showModal();
    });

    closeModal.addEventListener("click", () => {
      modal.close();
      form.clear();
    });

    modal.addEventListener("slot:assigned", () => modal.close());
  }
}
