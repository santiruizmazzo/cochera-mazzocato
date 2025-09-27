import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
  }

  async getHtml() {
    return /*html*/ `
      <section class="home-view">
        <h2>¡Bienvenido a la app de gestión de la <span>Cochera Mazzocato</span>!</h2>
      </section>
    `;
  }
}
