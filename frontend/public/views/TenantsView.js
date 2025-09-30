import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
  }

  getHtml() {
    return /*html*/ `
      <dialog id="new-tenant-modal" closedby="any">
        <header>
          <h3>Registrar inquilino</h3>
          <button class="close-action-btn close-modal-btn">
            <svg><use href="#close"/></svg>
          </button>
        </header>
        <tenant-form></tenant-form>
      </dialog>

      <header class="section-header">
        <h2>Inquilinos registrados</h2>
        <button class="new-action-btn open-modal-btn">
          <svg><use href="#person_add"/></svg>
          Nuevo inquilino
        </button>
      </header>
      
      <tenants-collection>
        <svg class="spinner"><use href="#spinner"/></svg>
      </tenants-collection>
    `;
  }

  setUpJavascript() {
    const modal = document.querySelector("#new-tenant-modal");
    const openModal = document.querySelector(".open-modal-btn");
    const closeModal = document.querySelector(".close-modal-btn");
    const form = document.querySelector("tenant-form");

    openModal.addEventListener("click", () => modal.showModal());
    closeModal.addEventListener("click", () => {
      modal.close();
      form.clear();
    });
    modal.addEventListener("tenants:update", () => modal.close());
  }
}
