import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("Inquilinos");
  }

  async getHtml() {
    return /*html*/ `
      <dialog class="modal" closedby="any">
        <tenant-form></tenant-form>
      </dialog>

      <header class="tenants-header">
        <h2 class="section-title">Inquilinos registrados</h2>
        <button class="open-modal">
          <span class="material-symbols-outlined">person_add</span>
          <p>Nuevo inquilino</p>
        </button>
      </header>
      
      <tenants-list></tenants-list>
    `;
  }

  setUpJavascript() {
    const modal = document.querySelector(".modal");
    const openModal = document.querySelector(".open-modal");

    openModal.addEventListener("click", () => {
      modal.showModal();
    });
  }
}
