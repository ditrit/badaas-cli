// Code generated by badaas-cli v0.0.0, DO NOT EDIT.
package belongsto

import orm "github.com/ditrit/badaas/orm"

func (m Owned) GetOwner() (*Owner, error) {
	return orm.VerifyStructLoaded[Owner](&m.Owner)
}
