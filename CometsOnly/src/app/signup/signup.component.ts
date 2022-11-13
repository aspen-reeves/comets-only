import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { Profile } from '../shared/profile.model';
import { SignupPost } from './signup.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css'],
})
export class SignupComponent implements OnInit {
  signup: FormGroup;

  constructor(private http: HttpClient, private signupPost: SignupPost) {}

  ngOnInit(): void {
    this.signup = new FormGroup({
      name: new FormControl('isaac', [Validators.required]),
      age: new FormControl(19, [Validators.required]),
      lang: new FormControl('js', [Validators.required]),
      os: new FormControl('windows', [Validators.required]),
      editor: new FormControl('vs code', [Validators.required]),
      lastShower: new FormControl('2022-11-12', [Validators.required]),
      code: new FormControl('console.log("Hello World")', [
        Validators.required,
      ]),
    });
    console.log(this.signup);
  }

  onSubmit(signupForm: Profile) {
    // console.log(signupForm);

    // let newProfile: Profile = { ...this.signup.value };
    this.signupPost.addProfile(signupForm);
  }

  // getBitches() {
  //   this.signupPost.getProfile();
  // }
}
