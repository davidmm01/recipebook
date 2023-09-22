REQUIRED_STRING = {"required": True, "type": "string"}
OPTIONAL_STRING = {"required": False, "type": "string"}
REQUIRED_DATE = {"required": True, "type": "date"}

SOURCE_TYPE_ALLOWED = [
    "copy",  # copied from the internet for easier referencing
]

DESCRIPTOR_TAGS_ALLOWED = ["condiment"]

CUISINE_ALLOWED = [
    "japanese",
]

SCHEMA = {
    "name": REQUIRED_STRING,
    "date_added": REQUIRED_DATE,
    "source": {
        "name": REQUIRED_STRING,
        "url": OPTIONAL_STRING,
        "type": {
            "required": False,
        },
    },
    "descriptor_tags": {
        "required": True,
        "type": "string",
        "allowed": DESCRIPTOR_TAGS_ALLOWED,
    },
    "ingredients": {
        "required": True,
        "type": "list",
        "items": [{"type": "string"}],
    },  # TODO: can i drop the array for items if just 1 element?
    "instructions": {"required": True, "type": "list", "items": [{"type": "string"}]},
    "usage": OPTIONAL_STRING,  # i.e. serving suggestions, if it's a condiment, what might it go with
}

# TODO: should we be doing a schema per type? doesn't make sense to have recipes, conditments, and maybe things like cocktails in the same schema!
# TODO: would be nice to have a special hardware tags so ppl can easily filter in or out things like bbq smoker, sous vide etc...
