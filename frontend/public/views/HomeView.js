import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  async getHtml() {
    return /*html*/ `
      <section class="home-view">
        <h2>¡Bienvenido a la app de gestión de la <span>Cochera Mazzocato</span>!</h2>
        
        <dialog id="assign-tenant-to-slot-modal" closedby="any">
          <header>
            <h3>Reservar plaza</h3>
            <button class="close-action-btn close-modal-btn">
              <svg><use href="#close"/></svg>
            </button>
          </header>
          <slot-form></slot-form>
        </dialog>

        <div class="slots-collection">
          <div id="1" class="slot">
            <h2>1</h2>
            <p>De <a href="/inquilinos/1" data-link>Lamponne</a></p>
          </div>
          <div id="2" class="slot">
            <h2>2</h2>
            <p>De <a href="/inquilinos/1" data-link>Lamponne</a></p>
          </div>
          <div id="3" class="slot">
            <h2>3</h2>
            <p>De <a href="/inquilinos/1" data-link>Lamponne</a></p>
          </div>
          <div id="4" class="slot">
            <h2>4</h2>
            <p>De <a href="/inquilinos/1" data-link>Lamponne</a></p>
          </div>
          <div id="5" class="slot">
            <h2>5</h2>
            <p>De <a href="/inquilinos/2" data-link>Ravenna</a></p>
          </div>
          <div id="6" class="slot">
            <h2>6</h2>
            <p>De <a href="/inquilinos/2" data-link>Ravenna</a></p>
          </div>
          <div id="7" class="slot">
            <h2>7</h2>
            <p>De <a href="/inquilinos/2" data-link>Ravenna</a></p>
          </div>
          <div id="8" class="slot">
            <h2>8</h2>
            <p>De <a href="/inquilinos/2" data-link>Ravenna</a></p>
          </div>
          <div id="9" class="slot">
            <h2>9</h2>
            <p>De <a href="/inquilinos/2" data-link>Ravenna</a></p>
          </div>
          <div id="10" class="slot">
            <h2>10</h2>
            <div class="assign-tenant-btn open-modal-btn">
              Asignar
              <svg><use href="#person_waving" /></svg>
            </div>
          </div>
          <div id="11" class="slot">
            <h2>11</h2>
            <div class="assign-tenant-btn open-modal-btn">
              Asignar
              <svg><use href="#person_waving" /></svg>
            </div>
          </div>
          <div id="12" class="slot">
            <h2>12</h2>
            <p>De <a href="/inquilinos/2" data-link>Ravenna</a></p>
          </div>
        </div>
        <div></div>
      </section>
    `;
  }

  setUpJavascript() {
    const modal = document.querySelector("#assign-tenant-to-slot-modal");
    const openModalButtons = document.querySelectorAll(".open-modal-btn");
    const closeModal = document.querySelector(".close-modal-btn");
    const form = document.querySelector("slot-form");

    openModalButtons.forEach((button) => {
      button.addEventListener("click", () => {
        form.slotNumber = button.parentElement.id;
        modal.showModal();
      });
    });
    closeModal.addEventListener("click", () => {
      modal.close();
      form.clear();
    });
  }
}
