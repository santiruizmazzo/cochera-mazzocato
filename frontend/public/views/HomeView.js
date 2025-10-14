import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  async getHtml() {
    return /*html*/ `
      <section class="home-view">
        <h2>¡Bienvenido a la app de gestión de la <span>Cochera Mazzocato</span>!</h2>
        <div class="slots-collection">
          <div class="slot">
            <h2>1</h2>
            <p>Alquilada por: <a href="/inquilinos/3" data-link>Ruiz Mazzocato</a></p>
          </div>
          <div class="slot">
            <h2>2</h2>
            <p>Alquilada por: <a href="/inquilinos/6" data-link>Lamponne</a></p>
          </div>
          <div class="slot">
            <h2>3</h2>
            <p>Alquilada por: <a href="/inquilinos/8" data-link>Ravenna</a></p>
          </div>
          <div class="slot">
            <h2>4</h2>
            <p>Alquilada por: <a href="/inquilinos/6" data-link>Lamponne</a></p>
          </div>
          <div class="slot">
            <h2>5</h2>
            <p>Alquilada por: <a href="/inquilinos/9" data-link>Santos</a></p>
          </div>
          <div class="slot">
            <h2>6</h2>
            <p>Alquilada por: <a href="/inquilinos/9" data-link>Santos</a></p>
          </div>
          <div class="slot">
            <h2>7</h2>
            <p>Alquilada por: <a href="/inquilinos/13" data-link>Hamlin</a></p>
          </div>
          <div class="slot">
            <h2>8</h2>
            <p>Alquilada por: <a href="/inquilinos/14" data-link>Goodman</a></p>
          </div>
          <div class="slot">
            <h2>9</h2>
            <p>Alquilada por: <a href="/inquilinos/10" data-link>Medina</a></p>
          </div>
          <div class="slot">
            <h2>10</h2>
            <p>Libre 😧</p>
          </div>
          <div class="slot">
            <h2>11</h2>
            <p>Libre 😧</p>
          </div>
          <div class="slot">
            <h2>12</h2>
            <p>Alquilada por: <a href="/inquilinos/14" data-link>Goodman</a></p>
          </div>
        </div>
      </section>
    `;
  }
}
