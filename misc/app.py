#!/usr/bin/env python3

from flask import Flask, jsonify
import json

app = Flask(__name__)
app.secret_key = "ps105"

emp_id_map = {
	'A123456789': 'Flynn Mackie',
	'A100000001': 'Wesley Thomas',
	'A100000002': 'Nina Chiswick',
	'A100000003': 'Randall Cosmo',
	'A100000004': 'Brenda Plager',
	'A100000005': 'Tommy Quinn',
	'A100000006': 'Jake Farmer',
	'A100000007': 'Liam Freeman',
	'A100000008': 'Sheila Dunbar',
	'A100000009': 'Peter Young'
}

emp_info = {
	'Flynn Mackie': {'title': 'Senior VP of Engineering', 'reports': ['A100000001','A100000002']},
	'Wesley Thomas': {'title': 'VP of Design', 'reports': ['A100000003']},
	'Nina Chiswick': {'title': 'VP of Engineering', 'reports': ['A100000005']},
	'Randall Cosmo': {'title': 'Director of Design', 'reports': ['A100000004']},
	'Brenda Plager': {'title': 'Senior Designer', 'reports': []},
	'Tommy Quinn': {'title': 'Director of Engineering', 'reports': ['A100000006','A100000008']},
	'Jake Farmer': {'title': 'Frontend Manager', 'reports': ['A100000007']},
	'Liam Freeman': {'title': 'Junior Code Monkey', 'reports': []},
	'Sheila Dunbar': {'title': 'Backend Manager', 'reports': ['A100000009']},
	'Peter Young': {'title': 'Senior Code Cowboy', 'reports': []}
}

@app.route('/linkedin/api/employee/<employee_id>', methods=['GET','POST'])    # For linked Api question
def employee_info(employee_id):
	name = emp_id_map[employee_id.upper()]
	final_dict = {'name': name}
	final_dict.update(emp_info[name])
	return jsonify(final_dict)

if __name__ == "__main__":
	app.run(port=37899)
