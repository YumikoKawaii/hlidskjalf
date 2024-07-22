def convert_str_to_camel(snake_str):
    parts = snake_str.split('_')
    return parts[0] + ''.join(x.title() for x in parts[1:])