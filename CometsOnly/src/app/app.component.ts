import { Component, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-root',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css', '../../node_modules/xp.css/dist/XP.css']
})
export class AppComponent {
  title = 'CometsOnly';

  profiles:number[] = [];

  

}
