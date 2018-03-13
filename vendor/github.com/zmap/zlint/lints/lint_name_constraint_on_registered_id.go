/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

/*******************************************************************
RFC 5280: 4.2.1.10
Restrictions are defined in terms of permitted or excluded name
subtrees.  Any name matching a restriction in the excludedSubtrees
field is invalid regardless of information appearing in the
permittedSubtrees.  Conforming CAs MUST mark this extension as
critical and SHOULD NOT impose name constraints on the x400Address,
ediPartyName, or registeredID name forms.  Conforming CAs MUST NOT
issue certificates where name constraints is an empty sequence.  That
is, either the permittedSubtrees field or the excludedSubtrees MUST
be present.
*******************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type nameConstraintOnRegisteredId struct{}

func (l *nameConstraintOnRegisteredId) Initialize() error {
	return nil
}

func (l *nameConstraintOnRegisteredId) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.NameConstOID)
}

func (l *nameConstraintOnRegisteredId) Execute(c *x509.Certificate) *LintResult {
	if c.PermittedRegisteredIDs != nil || c.ExcludedRegisteredIDs != nil {
		return &LintResult{Status: Warn}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_name_constraint_on_registered_id",
		Description:   "The name constraints extension SHOULD NOT impose constraints on the registeredID name form",
		Citation:      "RFC 5280: 4.2.1.10",
		Source:        RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          &nameConstraintOnRegisteredId{},
	})
}
