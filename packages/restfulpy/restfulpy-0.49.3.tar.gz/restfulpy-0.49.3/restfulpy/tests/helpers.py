from restfulpy.application import Application


# noinspection PyAbstractClass
class MockupApplication(Application):
    builtin_configuration = '''
    db:
      test_url: postgresql://postgres:postgres@localhost/restfulpy_test
      administrative_url: postgresql://postgres:postgres@localhost/postgres      
    logging:
      loggers:
        default:
          level: critical
    '''

    def configure(self, files=None, context=None, **kwargs):
        _context = dict(
            process_name='restfulpy_unittests'
        )
        if context:
            _context.update(context)
        super().configure(files=files, context=_context, **kwargs)
