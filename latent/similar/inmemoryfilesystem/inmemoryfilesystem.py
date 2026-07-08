class Node:
    def __init__(self, is_file):
        self.is_file = is_file
        self.content = []
        self.children = {}


class Filesystem:
    def __init__(self):
        self.root = Node(False)
        self.root.children[""] = Node(False)

    def _walk_to_node(self, components, create_path: bool) -> Node:
        current = self.root
        for i in range(len(components)):
            if components[i] not in current.children:
                if not create_path:
                    raise Exception("invalid path")
                current.children[components[i]] = Node(False)
            current = current.children[components[i]]
        return current

    def ls(self, filepath) -> list[str]:
        components = filepath.strip("/").split("/")
        node = self._walk_to_node(components, False)
        contents = [x for x in list(node.children.keys()) if x != ""]
        return contents

    def add_content_to_file(self, filepath: str, content: str) -> None:
        components = filepath.split("/")
        location = components  # exclude file name from path
        node = self._walk_to_node(location, True)

        encoded = list(content.encode("utf-8"))
        node.content.extend(encoded)


fs = Filesystem()
assert fs.ls("/") == []
assert fs.add_content_to_file("/hello.txt", "start") is None
assert fs.ls("/") == ["hello.txt"]
assert fs.add_content_to_file("/home/logs.txt", "somelogs") is None
assert fs.ls("/") == ["hello.txt", "home"]
