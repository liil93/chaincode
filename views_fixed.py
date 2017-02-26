from flask import render_template, flash, redirect, request, Response
from flask import Flask, session, url_for, escape
from app import app
import json, requests
from . import function
from . import blockchain_restapi
# set the secret key.  keep this really secret:
app.secret_key = 'A0Zr98j/3yX R~XHH!jmN]LWX/,?RT'
# app.secret_key = os.urandom(24)


# index view function suppressed for brevity
@app.route('/')
@app.route('/index')
def index():



	if 'email' in session:
		return render_template("search.html",
                        title='Welcome',
						session=session['email'])
	else:
		return render_template("search.html",
                        title='Welcome',
						session=None)


@app.route('/login', methods=['GET', 'POST'])
def login():
	if 'email' in session:
		return redirect('/')
	else:
	    error=None
	    if request.method == 'POST':
	        Email= request.form.get("email")
	        PW = request.form.get("password")
	        result1 = function.Check_email(Email)
	        result2 = function.Check_pw(Email, PW)
	        if result1 ==[]:
	            error = 'Invalid username'
	        elif result2 ==[]:
	            error = 'Invalid password'
	        else:
	            session['email'] = Email
	            return redirect('/')

	    return render_template("login.html",
	                        title='Sign In',
	                        error=error)

@app.route('/signup', methods=['POST', 'GET'])
def signup():
	# ImmutableMultiDict([('user[last_name]', 'yoo'),
	# 					('user[password]', '1q2w3e4r'),
	# 					('user[birthday_month]', '1'),
	# 					('user[birthday_day]', '4'),
	# 					('user[birthday_year]', '2011'),
	# 					('utf8', '✓'),
	# 					('from', 'email_signup'),
	# 					('user[first_name]', 'dongwon'),
	# 					('user[email]', 'asldkfj@nave.rcom'),
	# 					('user_profile_info[receive_promotional_email]', '0'),
	# 					('user_profile_info[receive_promotional_email]', '1'),
	# 					('authenticity_token', '#j')])

	if request.method == 'POST':
		Email= request.form.get("user[email]")
		PW = request.form.get("user[password]")
		result = function.Save_mem(Email, PW)
		if result == 1:
			session['email'] = Email
			return redirect('/')
		else:
			error = 'Email already exists.'
			return render_template("signup.html",
								title='SignUp',
								error=error)
	else:
		return render_template("signup.html",
							title='SignUp',
							error=None)





@app.route('/logout', methods=['GET', 'POST'])
def logout():
    # remove the username from the session if it's there
    session.pop('email', None)
    return redirect('/')


@app.route('/enrollment_home/address', methods=['GET', 'POST'])
def enrollment_home_address():
	# Check session
	if not 'email' in session:
		return redirect('/')

	if request.method == 'POST':
		print(request.form)
		# ImmutableMultiDict([('country_code', 'US'),
		# 					('city', '1'),
		# 					('street', '1'),
		# 					('zipcode', '1'),
		# 					('apt', '1'),
		# 					('state', '1')])
		return redirect('/enrollment_home/room')

	return render_template("address.html",
                        title='progress',
						session='OK')


@app.route('/enrollment_home/room', methods=['GET', 'POST'])
def enrollment_home_room():
	# Check session
	if not 'email' in session:
		return redirect('/')


	if request.method == 'POST':
		print(request.form)
		# ImmutableMultiDict([('house_type', '2'),
		# 					('number_of_room', '1')])
		return redirect('/enrollment_home/car_elevator')

	return render_template("room.html",
                        title='progress',
						session='OK')

@app.route('/enrollment_home/car_elevator', methods=['GET', 'POST'])
def enrollment_home_car_elevator():
	# Check session
	if not 'email' in session:
		return redirect('/')


	if request.method == 'POST':
		print(request.form)
		# ImmutableMultiDict([('elevatorType', 'yes'),
		# 					('parkingType', 'no')])
		return redirect('/enrollment_home/complete')

	return render_template("car_elevator.html",
                        title='progress',
						session='OK')

@app.route('/enrollment_home/complete', methods=['GET', 'POST'])
def enrollment_home_complete():
	# Check session
	if not 'email' in session:
		return redirect('/')


	if request.method == 'POST':
		print(request.form)
		# ImmutableMultiDict([('elevatorType', 'yes'),
		# 					('parkingType', 'no')])
		return redirect('/')

	return render_template("enrollment_home_complete.html",
                        title='progress',
						session='OK')

@app.route('/test_login', methods=['GET', 'POST'])
def test_login():
	return blockchain_restapi.login()

@app.route('/test_init', methods=['GET', 'POST'])
def test_init():
	return blockchain_restapi.init()

@app.route('/test_user_insert', methods=['GET', 'POST'])
def test_user_insert():
	return blockchain_restapi.user_insert('liil93', 'qqwweerr')

@app.route('/test_home_insert', methods=['GET', 'POST'])
def test_home_insert():
	return blockchain_restapi.home_insert('liil93', 'R103', 'seoul', 'zachi', '3room','10m^2', 'Y', 'N')

@app.route('/test_pet_insert', methods=['GET', 'POST'])
def test_pet_insert():
	return blockchain_restapi.pet_insert('liil93', 'ddog', '170225', 'male', 'ddongdog', '1kg', 'Y', 'N')

@app.route('/test_user_change', methods=['GET', 'POST'])
def test_user_change():
	return blockchain_restapi.user_change('liil93', 'newqqwweerr', '1')

@app.route('/test_pet_change', methods=['GET', 'POST'])
def test_pet_change():
	return blockchain_restapi.pet_change('liil93', '20kg', 'Y', 'Y')

@app.route('/test_home_delete', methods=['GET', 'POST'])
def test_home_delete():
	return blockchain_restapi.home_delete('liil93')

@app.route('/test_pet_delete', methods=['GET', 'POST'])
def test_pet_delete():
	return blockchain_restapi.pet_delete('liil93')

@app.route('/test_trade_insert', methods=['GET', 'POST'])
def test_trade_insert():
	return blockchain_restapi.trade_insert('liil93', 'sanhak722', '931228', '201228', '121228', '500$', 'good')

@app.route('/test_user_read', methods=['GET', 'POST'])
def test_user_read():
	return blockchain_restapi.user_read('liil93')

@app.route('/test_home_read', methods=['GET', 'POST'])
def test_home_read():
	return blockchain_restapi.home_read('liil93')

@app.route('/test_pet_read', methods=['GET', 'POST'])
def test_pet_read():
	return blockchain_restapi.pet_read('liil93')

@app.route('/test_city_search', methods=['GET', 'POST'])
def test_city_search():
	return blockchain_restapi.city_search('R103')

@app.route('/test_trade_search', methods=['GET', 'POST'])
def test_trade_search():
	return blockchain_restapi.trade_search('liil93', 'sanhak722', '121228')
