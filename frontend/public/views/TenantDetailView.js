import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  async getHtml() {
    return /*html*/ `
      <section class="tenant-detail-view">
        <h2>Aca deberia mostrarte la info del inquilino con id = ${this.params.id}</h2>
      </section>
    `;
  }
}
