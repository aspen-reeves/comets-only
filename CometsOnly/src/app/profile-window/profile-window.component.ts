import { Component, Input, OnInit } from '@angular/core';
import { Profile } from '../shared/profile.model';

@Component({
  selector: 'app-profile-window',
  templateUrl: './profile-window.component.html',
  styleUrls: ['./profile-window.component.css'],
})
export class ProfileWindowComponent implements OnInit {
  @Input() profile: Profile;

  constructor() {}

  ngOnInit(): void {}


}
