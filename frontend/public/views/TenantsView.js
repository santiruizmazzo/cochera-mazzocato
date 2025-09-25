import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("Inquilinos");
  }

  getHtml() {
    return /*html*/ `
      <dialog class="modal" closedby="any">
        <h3>Registrar inquilino</h3>
        <button class="close-modal">
          <svg><use href="assets/icons.svg#close"/></svg>
        </button>
        <tenant-form></tenant-form>
      </dialog>

      <header class="tenants-header">
        <h2 class="section-title">Inquilinos registrados</h2>
        <button class="open-modal">
          <svg><use href="assets/icons.svg#person_add"/></svg>
          <p>Nuevo inquilino</p>
        </button>
      </header>
      
      <tenants-list></tenants-list>
    `;
  }

  setUpJavascript() {
    const modal = document.querySelector(".modal");
    const openModal = document.querySelector(".open-modal");
    const closeModal = document.querySelector(".close-modal");

    openModal.addEventListener("click", () => {
      modal.showModal();
    });

    closeModal.addEventListener("click", () => {
      modal.close();
    });
  }
}
