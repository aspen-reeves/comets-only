import { Component, OnInit, Output } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { EventEmitter } from 'events';
import { APIService } from '../shared/api.service';

@Component({
  selector: 'app-signin',
  templateUrl: './signin.component.html',
  styleUrls: ['./signin.component.css'],
})
export class SigninComponent implements OnInit {
  signIn: FormGroup;
  @Output() signingUp: EventEmitter = new EventEmitter();
  signedIn = false;
  submitted = false;

  constructor(private API: APIService, public router: Router) {}

  ngOnInit(): void {
    this.signIn = new FormGroup({
      username: new FormControl(null, [Validators.required]),
      password: new FormControl(null, [Validators.required]),
    });
  }

  login(credentials: { u: string; p: string }) {
    this.submitted = true;

    this.API.signin(credentials);
    if (this.API.signedIn) {
      this.signedIn = true;
      this.router.navigate(['signup']);
    } else {
      this.signedIn = false;
    }
  }
}
