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
    // CREATE SIGNUP FORM
    this.signup = new FormGroup({
      name: new FormControl('name', [Validators.required]),
      age: new FormControl(18, [Validators.required]),
      lang: new FormControl('C++', [Validators.required]),
      os: new FormControl('Windows', [Validators.required]),
      editor: new FormControl('Visual Studio', [Validators.required]),
      lastShower: new FormControl('2022-11-12', [Validators.required]),
      code: new FormControl('cout << "Hello World!"', [
        Validators.required,
      ]),
    });
    console.log(this.signup);
  }

  // POST REQUEST TO SIGNUP
  onSubmit(signupForm: Profile) {
    this.signupPost.addProfile(signupForm);
  }
}
