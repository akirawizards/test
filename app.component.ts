import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <h1>Hello, Angular!</h1>
    <p>This repository is detected as TypeScript.</p>
  `,
  styles: [`
    h1 { color: #dd0031; }
  `]
})
export class AppComponent {
  title = 'angular-detection-sample';

  constructor() {
    console.log('Angular app is running!');
  }
}
