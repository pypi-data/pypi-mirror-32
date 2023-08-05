
# addr_spec_parsetab.py
# This file is automatically generated. Do not edit.
_tabversion = '3.10'

_lr_method = 'LALR'

_lr_signature = 'addr_specFWSP AT DOT COMMA SEMICOLON LANGLE RANGLE ATOM DOT_ATOM LBRACKET RBRACKET DTEXT DQUOTE QTEXT QPAIR LPAREN RPAREN CTEXT URLmailbox_or_url_list : mailbox_or_url_list delim mailbox_or_url\n                           | mailbox_or_url_list delim\n                           | mailbox_or_urldelim : delim fwsp COMMA\n             | delim fwsp SEMICOLON\n             | COMMA\n             | SEMICOLONmailbox_or_url : mailbox\n                      | urlurl : ofwsp URL ofwspmailbox : addr_spec\n               | angle_addr\n               | name_addrname_addr : ofwsp phrase angle_addrangle_addr : ofwsp LANGLE addr_spec RANGLE ofwspaddr_spec : ofwsp local_part AT domain ofwsplocal_part : DOT_ATOM\n                  | ATOM\n                  | quoted_stringdomain : DOT_ATOM\n              | ATOM\n              | domain_literalquoted_string : DQUOTE quoted_string_text DQUOTE\n                     | DQUOTE DQUOTEquoted_string_text : quoted_string_text QTEXT\n                          | quoted_string_text QPAIR\n                          | quoted_string_text fwsp\n                          | QTEXT\n                          | QPAIR\n                          | fwspdomain_literal : LBRACKET domain_literal_text RBRACKET\n                      | LBRACKET RBRACKETdomain_literal_text : domain_literal_text DTEXT\n                           | domain_literal_text fwsp\n                           | DTEXT\n                           | fwspcomment : LPAREN comment_text RPAREN\n               | LPAREN RPARENcomment_text : comment_text CTEXT\n                    | comment_text fwsp\n                    | CTEXT\n                    | fwspphrase : phrase fwsp ATOM\n              | phrase fwsp DOT_ATOM\n              | phrase fwsp DOT\n              | phrase fwsp quoted_string\n              | phrase ATOM\n              | phrase DOT_ATOM\n              | phrase DOT\n              | phrase quoted_string\n              | ATOM\n              | DOT_ATOM\n              | DOT\n              | quoted_stringofwsp : fwsp comment fwsp\n             | fwsp comment\n             | comment fwsp\n             | comment\n             | fwsp\n             |fwsp : FWSP'
    
_lr_action_items = {'FWSP':([0,2,4,6,7,10,14,15,16,17,20,21,22,24,25,26,27,28,29,30,31,32,33,34,35,38,39,40,41,42,43,44,],[4,4,-61,4,4,4,-42,4,-38,-41,-30,-29,-28,4,-40,-37,-39,4,4,-20,-21,-22,-27,-26,-25,-36,-35,4,-32,-34,-33,-31,]),'QPAIR':([4,10,20,21,22,24,33,34,35,],[-61,21,-30,-29,-28,34,-27,-26,-25,]),'QTEXT':([4,10,20,21,22,24,33,34,35,],[-61,22,-30,-29,-28,35,-27,-26,-25,]),'DTEXT':([4,29,38,39,40,42,43,],[-61,39,-36,-35,43,-34,-33,]),'DQUOTE':([0,1,2,3,4,7,8,10,16,18,20,21,22,24,26,33,34,35,],[-60,-59,-58,10,-61,-56,-57,23,-38,-55,-30,-29,-28,36,-37,-27,-26,-25,]),'LBRACKET':([19,],[29,]),'DOT_ATOM':([0,1,2,3,4,7,8,16,18,19,26,],[-60,-59,-58,11,-61,-56,-57,-38,-55,30,-37,]),'AT':([9,11,12,13,23,36,],[19,-17,-18,-19,-24,-23,]),'LPAREN':([0,1,4,28,30,31,32,41,44,],[6,6,-61,6,-20,-21,-22,-32,-31,]),'ATOM':([0,1,2,3,4,7,8,16,18,19,26,],[-60,-59,-58,12,-61,-56,-57,-38,-55,31,-37,]),'RPAREN':([4,6,14,15,17,25,27,],[-61,16,-42,26,-41,-40,-39,]),'RBRACKET':([4,29,38,39,40,42,43,],[-61,41,-36,-35,44,-34,-33,]),'CTEXT':([4,6,14,15,17,25,27,],[-61,17,-42,27,-41,-40,-39,]),'$end':([1,2,4,5,7,8,16,18,26,28,30,31,32,37,41,44,],[-59,-58,-61,0,-56,-57,-38,-55,-37,-60,-20,-21,-22,-16,-32,-31,]),}

_lr_action = {}
for _k, _v in _lr_action_items.items():
   for _x,_y in zip(_v[0],_v[1]):
      if not _x in _lr_action:  _lr_action[_x] = {}
      _lr_action[_x][_k] = _y
del _lr_action_items

_lr_goto_items = {'fwsp':([0,2,6,7,10,15,24,28,29,40,],[1,8,14,18,20,25,33,1,38,42,]),'comment':([0,1,28,],[2,7,2,]),'domain':([19,],[28,]),'comment_text':([6,],[15,]),'ofwsp':([0,28,],[3,37,]),'local_part':([3,],[9,]),'quoted_string_text':([10,],[24,]),'domain_literal_text':([29,],[40,]),'addr_spec':([0,],[5,]),'quoted_string':([3,],[13,]),'domain_literal':([19,],[32,]),}

_lr_goto = {}
for _k, _v in _lr_goto_items.items():
   for _x, _y in zip(_v[0], _v[1]):
       if not _x in _lr_goto: _lr_goto[_x] = {}
       _lr_goto[_x][_k] = _y
del _lr_goto_items
_lr_productions = [
  ("S' -> addr_spec","S'",1,None,None,None),
  ('mailbox_or_url_list -> mailbox_or_url_list delim mailbox_or_url','mailbox_or_url_list',3,'p_expression_mailbox_or_url_list','parser.py',19),
  ('mailbox_or_url_list -> mailbox_or_url_list delim','mailbox_or_url_list',2,'p_expression_mailbox_or_url_list','parser.py',20),
  ('mailbox_or_url_list -> mailbox_or_url','mailbox_or_url_list',1,'p_expression_mailbox_or_url_list','parser.py',21),
  ('delim -> delim fwsp COMMA','delim',3,'p_delim','parser.py',30),
  ('delim -> delim fwsp SEMICOLON','delim',3,'p_delim','parser.py',31),
  ('delim -> COMMA','delim',1,'p_delim','parser.py',32),
  ('delim -> SEMICOLON','delim',1,'p_delim','parser.py',33),
  ('mailbox_or_url -> mailbox','mailbox_or_url',1,'p_expression_mailbox_or_url','parser.py',36),
  ('mailbox_or_url -> url','mailbox_or_url',1,'p_expression_mailbox_or_url','parser.py',37),
  ('url -> ofwsp URL ofwsp','url',3,'p_expression_url','parser.py',41),
  ('mailbox -> addr_spec','mailbox',1,'p_expression_mailbox','parser.py',45),
  ('mailbox -> angle_addr','mailbox',1,'p_expression_mailbox','parser.py',46),
  ('mailbox -> name_addr','mailbox',1,'p_expression_mailbox','parser.py',47),
  ('name_addr -> ofwsp phrase angle_addr','name_addr',3,'p_expression_name_addr','parser.py',51),
  ('angle_addr -> ofwsp LANGLE addr_spec RANGLE ofwsp','angle_addr',5,'p_expression_angle_addr','parser.py',55),
  ('addr_spec -> ofwsp local_part AT domain ofwsp','addr_spec',5,'p_expression_addr_spec','parser.py',59),
  ('local_part -> DOT_ATOM','local_part',1,'p_expression_local_part','parser.py',63),
  ('local_part -> ATOM','local_part',1,'p_expression_local_part','parser.py',64),
  ('local_part -> quoted_string','local_part',1,'p_expression_local_part','parser.py',65),
  ('domain -> DOT_ATOM','domain',1,'p_expression_domain','parser.py',69),
  ('domain -> ATOM','domain',1,'p_expression_domain','parser.py',70),
  ('domain -> domain_literal','domain',1,'p_expression_domain','parser.py',71),
  ('quoted_string -> DQUOTE quoted_string_text DQUOTE','quoted_string',3,'p_expression_quoted_string','parser.py',75),
  ('quoted_string -> DQUOTE DQUOTE','quoted_string',2,'p_expression_quoted_string','parser.py',76),
  ('quoted_string_text -> quoted_string_text QTEXT','quoted_string_text',2,'p_expression_quoted_string_text','parser.py',83),
  ('quoted_string_text -> quoted_string_text QPAIR','quoted_string_text',2,'p_expression_quoted_string_text','parser.py',84),
  ('quoted_string_text -> quoted_string_text fwsp','quoted_string_text',2,'p_expression_quoted_string_text','parser.py',85),
  ('quoted_string_text -> QTEXT','quoted_string_text',1,'p_expression_quoted_string_text','parser.py',86),
  ('quoted_string_text -> QPAIR','quoted_string_text',1,'p_expression_quoted_string_text','parser.py',87),
  ('quoted_string_text -> fwsp','quoted_string_text',1,'p_expression_quoted_string_text','parser.py',88),
  ('domain_literal -> LBRACKET domain_literal_text RBRACKET','domain_literal',3,'p_expression_domain_literal','parser.py',92),
  ('domain_literal -> LBRACKET RBRACKET','domain_literal',2,'p_expression_domain_literal','parser.py',93),
  ('domain_literal_text -> domain_literal_text DTEXT','domain_literal_text',2,'p_expression_domain_literal_text','parser.py',100),
  ('domain_literal_text -> domain_literal_text fwsp','domain_literal_text',2,'p_expression_domain_literal_text','parser.py',101),
  ('domain_literal_text -> DTEXT','domain_literal_text',1,'p_expression_domain_literal_text','parser.py',102),
  ('domain_literal_text -> fwsp','domain_literal_text',1,'p_expression_domain_literal_text','parser.py',103),
  ('comment -> LPAREN comment_text RPAREN','comment',3,'p_expression_comment','parser.py',107),
  ('comment -> LPAREN RPAREN','comment',2,'p_expression_comment','parser.py',108),
  ('comment_text -> comment_text CTEXT','comment_text',2,'p_expression_comment_text','parser.py',112),
  ('comment_text -> comment_text fwsp','comment_text',2,'p_expression_comment_text','parser.py',113),
  ('comment_text -> CTEXT','comment_text',1,'p_expression_comment_text','parser.py',114),
  ('comment_text -> fwsp','comment_text',1,'p_expression_comment_text','parser.py',115),
  ('phrase -> phrase fwsp ATOM','phrase',3,'p_expression_phrase','parser.py',119),
  ('phrase -> phrase fwsp DOT_ATOM','phrase',3,'p_expression_phrase','parser.py',120),
  ('phrase -> phrase fwsp DOT','phrase',3,'p_expression_phrase','parser.py',121),
  ('phrase -> phrase fwsp quoted_string','phrase',3,'p_expression_phrase','parser.py',122),
  ('phrase -> phrase ATOM','phrase',2,'p_expression_phrase','parser.py',123),
  ('phrase -> phrase DOT_ATOM','phrase',2,'p_expression_phrase','parser.py',124),
  ('phrase -> phrase DOT','phrase',2,'p_expression_phrase','parser.py',125),
  ('phrase -> phrase quoted_string','phrase',2,'p_expression_phrase','parser.py',126),
  ('phrase -> ATOM','phrase',1,'p_expression_phrase','parser.py',127),
  ('phrase -> DOT_ATOM','phrase',1,'p_expression_phrase','parser.py',128),
  ('phrase -> DOT','phrase',1,'p_expression_phrase','parser.py',129),
  ('phrase -> quoted_string','phrase',1,'p_expression_phrase','parser.py',130),
  ('ofwsp -> fwsp comment fwsp','ofwsp',3,'p_expression_ofwsp','parser.py',139),
  ('ofwsp -> fwsp comment','ofwsp',2,'p_expression_ofwsp','parser.py',140),
  ('ofwsp -> comment fwsp','ofwsp',2,'p_expression_ofwsp','parser.py',141),
  ('ofwsp -> comment','ofwsp',1,'p_expression_ofwsp','parser.py',142),
  ('ofwsp -> fwsp','ofwsp',1,'p_expression_ofwsp','parser.py',143),
  ('ofwsp -> <empty>','ofwsp',0,'p_expression_ofwsp','parser.py',144),
  ('fwsp -> FWSP','fwsp',1,'p_expression_fwsp','parser.py',148),
]
