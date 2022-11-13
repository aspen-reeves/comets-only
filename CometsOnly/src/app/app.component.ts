import { Component, ViewEncapsulation } from '@angular/core';
import { APIService } from './shared/api.service';
import { Profile } from './shared/profile.model';

@Component({
  selector: 'app-root',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css', '../../node_modules/xp.css/dist/XP.css'],
})
export class AppComponent {
  title = 'CometsOnly';

  profiles: Profile[] = [];

  constructor(private API: APIService) {}

  // GETS A MATCH PROFILE
  match(): Profile {
    this.API.getProfile();
    return this.API.profile;
  }
}
