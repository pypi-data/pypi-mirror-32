class SingletonMeta(type):
    def __new__(mcs, name, bases, dct):
        ist = type.__new__(mcs, name, bases, dct)
        ist.__instance__ = None
        return ist

    def __call__(cls, *args, **kwargs):
        if cls.__instance__ is None:
            cls.__instance__ = type.__call__(cls, *args, **kwargs)
        return cls.__instance__


class Singleton:
    def __init__(self, clazz):
        self.clazz = clazz
        self.instance = None

    def __call__(self, *args, **kwargs):
        if self.instance is None:
            self.instance = self.clazz(*args, **kwargs)
            self.instance.__class__ = self.clazz
        return self.instance

    def __getattr__(self, item):
        return getattr(self.clazz, item)

    def __instancecheck__(self, instance):
        return isinstance(instance, self.clazz)
