// Code generated by triggergen. DO NOT EDIT.

package model

import (
	"fmt"

	"github.com/juju/juju/core/database/schema"
)


// ChangeLogTriggersForSubnet generates the triggers for the 
// subnet table.
func ChangeLogTriggersForSubnet(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for Subnet
CREATE TRIGGER trg_log_subnet_insert
AFTER INSERT ON subnet FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for Subnet
CREATE TRIGGER trg_log_subnet_update
AFTER UPDATE ON subnet FOR EACH ROW
WHEN 
	NEW.cidr != OLD.cidr OR
	(NEW.vlan_tag != OLD.vlan_tag OR (NEW.vlan_tag IS NOT NULL AND OLD.vlan_tag IS NULL) OR (NEW.vlan_tag IS NULL AND OLD.vlan_tag IS NOT NULL)) OR
	(NEW.space_uuid != OLD.space_uuid OR (NEW.space_uuid IS NOT NULL AND OLD.space_uuid IS NULL) OR (NEW.space_uuid IS NULL AND OLD.space_uuid IS NOT NULL)) 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for Subnet
CREATE TRIGGER trg_log_subnet_delete
AFTER DELETE ON subnet FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}
