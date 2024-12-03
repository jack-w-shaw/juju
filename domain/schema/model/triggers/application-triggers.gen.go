// Code generated by triggergen. DO NOT EDIT.

package triggers

import (
	"fmt"

	"github.com/juju/juju/core/database/schema"
)


// ChangeLogTriggersForApplication generates the triggers for the
// application table.
func ChangeLogTriggersForApplication(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert namespace for Application
INSERT INTO change_log_namespace VALUES (%[2]d, 'application', 'Application changes based on %[1]s');

-- insert trigger for Application
CREATE TRIGGER trg_log_application_insert
AFTER INSERT ON application FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for Application
CREATE TRIGGER trg_log_application_update
AFTER UPDATE ON application FOR EACH ROW
WHEN 
	NEW.uuid != OLD.uuid OR
	NEW.name != OLD.name OR
	NEW.life_id != OLD.life_id OR
	NEW.charm_uuid != OLD.charm_uuid OR
	(NEW.charm_modified_version != OLD.charm_modified_version OR (NEW.charm_modified_version IS NOT NULL AND OLD.charm_modified_version IS NULL) OR (NEW.charm_modified_version IS NULL AND OLD.charm_modified_version IS NOT NULL)) OR
	(NEW.charm_upgrade_on_error != OLD.charm_upgrade_on_error OR (NEW.charm_upgrade_on_error IS NOT NULL AND OLD.charm_upgrade_on_error IS NULL) OR (NEW.charm_upgrade_on_error IS NULL AND OLD.charm_upgrade_on_error IS NOT NULL)) OR
	(NEW.exposed != OLD.exposed OR (NEW.exposed IS NOT NULL AND OLD.exposed IS NULL) OR (NEW.exposed IS NULL AND OLD.exposed IS NOT NULL)) OR
	(NEW.placement != OLD.placement OR (NEW.placement IS NOT NULL AND OLD.placement IS NULL) OR (NEW.placement IS NULL AND OLD.placement IS NOT NULL)) OR
	(NEW.password_hash_algorithm_id != OLD.password_hash_algorithm_id OR (NEW.password_hash_algorithm_id IS NOT NULL AND OLD.password_hash_algorithm_id IS NULL) OR (NEW.password_hash_algorithm_id IS NULL AND OLD.password_hash_algorithm_id IS NOT NULL)) OR
	(NEW.password_hash != OLD.password_hash OR (NEW.password_hash IS NOT NULL AND OLD.password_hash IS NULL) OR (NEW.password_hash IS NULL AND OLD.password_hash IS NOT NULL)) 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;
-- delete trigger for Application
CREATE TRIGGER trg_log_application_delete
AFTER DELETE ON application FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

// ChangeLogTriggersForApplicationScale generates the triggers for the
// application_scale table.
func ChangeLogTriggersForApplicationScale(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert namespace for ApplicationScale
INSERT INTO change_log_namespace VALUES (%[2]d, 'application_scale', 'ApplicationScale changes based on %[1]s');

-- insert trigger for ApplicationScale
CREATE TRIGGER trg_log_application_scale_insert
AFTER INSERT ON application_scale FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for ApplicationScale
CREATE TRIGGER trg_log_application_scale_update
AFTER UPDATE ON application_scale FOR EACH ROW
WHEN 
	NEW.application_uuid != OLD.application_uuid OR
	(NEW.scale != OLD.scale OR (NEW.scale IS NOT NULL AND OLD.scale IS NULL) OR (NEW.scale IS NULL AND OLD.scale IS NOT NULL)) OR
	(NEW.scale_target != OLD.scale_target OR (NEW.scale_target IS NOT NULL AND OLD.scale_target IS NULL) OR (NEW.scale_target IS NULL AND OLD.scale_target IS NOT NULL)) OR
	(NEW.scaling != OLD.scaling OR (NEW.scaling IS NOT NULL AND OLD.scaling IS NULL) OR (NEW.scaling IS NULL AND OLD.scaling IS NOT NULL)) 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;
-- delete trigger for ApplicationScale
CREATE TRIGGER trg_log_application_scale_delete
AFTER DELETE ON application_scale FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

// ChangeLogTriggersForCharm generates the triggers for the
// charm table.
func ChangeLogTriggersForCharm(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert namespace for Charm
INSERT INTO change_log_namespace VALUES (%[2]d, 'charm', 'Charm changes based on %[1]s');

-- insert trigger for Charm
CREATE TRIGGER trg_log_charm_insert
AFTER INSERT ON charm FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for Charm
CREATE TRIGGER trg_log_charm_update
AFTER UPDATE ON charm FOR EACH ROW
WHEN 
	NEW.uuid != OLD.uuid OR
	(NEW.archive_path != OLD.archive_path OR (NEW.archive_path IS NOT NULL AND OLD.archive_path IS NULL) OR (NEW.archive_path IS NULL AND OLD.archive_path IS NOT NULL)) OR
	(NEW.available != OLD.available OR (NEW.available IS NOT NULL AND OLD.available IS NULL) OR (NEW.available IS NULL AND OLD.available IS NOT NULL)) OR
	(NEW.charmhub_identifier != OLD.charmhub_identifier OR (NEW.charmhub_identifier IS NOT NULL AND OLD.charmhub_identifier IS NULL) OR (NEW.charmhub_identifier IS NULL AND OLD.charmhub_identifier IS NOT NULL)) OR
	(NEW.version != OLD.version OR (NEW.version IS NOT NULL AND OLD.version IS NULL) OR (NEW.version IS NULL AND OLD.version IS NOT NULL)) OR
	NEW.source_id != OLD.source_id OR
	NEW.revision != OLD.revision OR
	(NEW.architecture_id != OLD.architecture_id OR (NEW.architecture_id IS NOT NULL AND OLD.architecture_id IS NULL) OR (NEW.architecture_id IS NULL AND OLD.architecture_id IS NOT NULL)) OR
	NEW.reference_name != OLD.reference_name 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;
-- delete trigger for Charm
CREATE TRIGGER trg_log_charm_delete
AFTER DELETE ON charm FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

// ChangeLogTriggersForPortRange generates the triggers for the
// port_range table.
func ChangeLogTriggersForPortRange(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert namespace for PortRange
INSERT INTO change_log_namespace VALUES (%[2]d, 'port_range', 'PortRange changes based on %[1]s');

-- insert trigger for PortRange
CREATE TRIGGER trg_log_port_range_insert
AFTER INSERT ON port_range FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for PortRange
CREATE TRIGGER trg_log_port_range_update
AFTER UPDATE ON port_range FOR EACH ROW
WHEN 
	NEW.uuid != OLD.uuid OR
	NEW.protocol_id != OLD.protocol_id OR
	(NEW.from_port != OLD.from_port OR (NEW.from_port IS NOT NULL AND OLD.from_port IS NULL) OR (NEW.from_port IS NULL AND OLD.from_port IS NOT NULL)) OR
	(NEW.to_port != OLD.to_port OR (NEW.to_port IS NOT NULL AND OLD.to_port IS NULL) OR (NEW.to_port IS NULL AND OLD.to_port IS NOT NULL)) OR
	(NEW.relation_uuid != OLD.relation_uuid OR (NEW.relation_uuid IS NOT NULL AND OLD.relation_uuid IS NULL) OR (NEW.relation_uuid IS NULL AND OLD.relation_uuid IS NOT NULL)) OR
	NEW.unit_uuid != OLD.unit_uuid 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;
-- delete trigger for PortRange
CREATE TRIGGER trg_log_port_range_delete
AFTER DELETE ON port_range FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

// ChangeLogTriggersForUnit generates the triggers for the
// unit table.
func ChangeLogTriggersForUnit(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert namespace for Unit
INSERT INTO change_log_namespace VALUES (%[2]d, 'unit', 'Unit changes based on %[1]s');

-- insert trigger for Unit
CREATE TRIGGER trg_log_unit_insert
AFTER INSERT ON unit FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for Unit
CREATE TRIGGER trg_log_unit_update
AFTER UPDATE ON unit FOR EACH ROW
WHEN 
	NEW.uuid != OLD.uuid OR
	NEW.name != OLD.name OR
	NEW.life_id != OLD.life_id OR
	NEW.application_uuid != OLD.application_uuid OR
	NEW.net_node_uuid != OLD.net_node_uuid OR
	(NEW.charm_uuid != OLD.charm_uuid OR (NEW.charm_uuid IS NOT NULL AND OLD.charm_uuid IS NULL) OR (NEW.charm_uuid IS NULL AND OLD.charm_uuid IS NOT NULL)) OR
	NEW.resolve_kind_id != OLD.resolve_kind_id OR
	(NEW.password_hash_algorithm_id != OLD.password_hash_algorithm_id OR (NEW.password_hash_algorithm_id IS NOT NULL AND OLD.password_hash_algorithm_id IS NULL) OR (NEW.password_hash_algorithm_id IS NULL AND OLD.password_hash_algorithm_id IS NOT NULL)) OR
	(NEW.password_hash != OLD.password_hash OR (NEW.password_hash IS NOT NULL AND OLD.password_hash IS NULL) OR (NEW.password_hash IS NULL AND OLD.password_hash IS NOT NULL)) 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;
-- delete trigger for Unit
CREATE TRIGGER trg_log_unit_delete
AFTER DELETE ON unit FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}
