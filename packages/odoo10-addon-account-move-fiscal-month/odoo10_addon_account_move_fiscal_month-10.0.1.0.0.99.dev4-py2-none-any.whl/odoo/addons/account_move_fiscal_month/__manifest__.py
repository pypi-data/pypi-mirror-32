# -*- coding: utf-8 -*-
# Copyright 2017 ACSONE SA/NV
# License AGPL-3.0 or later (http://www.gnu.org/licenses/agpl).

{
    'name': 'Account Move Fiscal Month',
    'summary': """
        Display the fiscal month on journal entries/item""",
    'version': '10.0.1.0.0',
    'license': 'AGPL-3',
    'author': 'ACSONE SA/NV,Odoo Community Association (OCA)',
    'website': 'https://www.acsone.eu',
    'depends': [
        'account_fiscal_month',
    ],
    'data': [
        'views/account_move.xml',
        'views/account_move_line.xml',
    ],
    'demo': [
    ],
}
