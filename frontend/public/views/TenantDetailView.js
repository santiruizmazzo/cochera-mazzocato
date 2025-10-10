import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  async getHtml() {
    return /*html*/ `
      <section class="tenant-detail-view">
        <tenant-full-card id="${this.params.id}"></tenant-full-card>
      </section>
    `;
  }
}
