import { Component, OnInit } from '@angular/core';
import { APIService } from '../shared/api.service';
import { Profile } from '../shared/profile.model';
import { Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-main-menu',
  templateUrl: './main-menu.component.html',
  styleUrls: ['./main-menu.component.css'],
})
export class MainMenuComponent implements OnInit {
  @Output() newProfileEvent = new EventEmitter<Profile>();
  profile: Profile;

  constructor(private API: APIService) {}

  ngOnInit(): void {
    // Continuously update 
    this.API.subject$.subscribe((data: Profile) => {
      this.profile = data;
    });
  }

  find() {
    this.API.getProfile();
  }
}
