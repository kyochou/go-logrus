package logrus

func (entry *Entry) ELog(level Level, err error, args ...interface{}) {
	if err != nil {
		entry.Log(level, args...)
		entry.Log(level, err)
	}
}

func (entry *Entry) ELogln(level Level, err error, args ...interface{}) {
	if err != nil {
		entry.Logln(level, args...)
		entry.Logln(level, err)
	}
}

func (entry *Entry) ELogf(level Level, err error, format string, args ...interface{}) {
	if err != nil {
		entry.Logf(level, format, args...)
		entry.Logf(level, err.Error())
	}
}

func (entry *Entry) EError(err error, args ...interface{}) {
	entry.ELog(ErrorLevel, err, args...)
}
func (entry *Entry) EErrorln(err error, args ...interface{}) {
	entry.ELogln(ErrorLevel, err, args...)
}
func (entry *Entry) EErrorf(err error, format string, args ...interface{}) {
	entry.ELogf(ErrorLevel, err, format, args...)
}

func (entry *Entry) EWarn(err error, args ...interface{}) {
	entry.ELog(WarnLevel, err, args...)
}
func (entry *Entry) EWarnln(err error, args ...interface{}) {
	entry.ELogln(WarnLevel, err, args...)
}
func (entry *Entry) EWarnf(err error, format string, args ...interface{}) {
	entry.ELogf(WarnLevel, err, format, args...)
}

func (entry *Entry) EInfo(err error, args ...interface{}) {
	entry.ELog(InfoLevel, err, args...)
}
func (entry *Entry) EInfoln(err error, args ...interface{}) {
	entry.ELogln(InfoLevel, err, args...)
}
func (entry *Entry) EInfof(err error, format string, args ...interface{}) {
	entry.ELogf(InfoLevel, err, format, args...)
}

func (entry *Entry) EDebug(err error, args ...interface{}) {
	entry.ELog(DebugLevel, err, args...)
}
func (entry *Entry) EDebugln(err error, args ...interface{}) {
	entry.ELogln(DebugLevel, err, args...)
}
func (entry *Entry) EDebugf(err error, format string, args ...interface{}) {
	entry.ELogf(DebugLevel, err, format, args...)
}

func (entry *Entry) ETrace(err error, args ...interface{}) {
	entry.ELog(TraceLevel, err, args...)
}
func (entry *Entry) ETraceln(err error, args ...interface{}) {
	entry.ELogln(TraceLevel, err, args...)
}
func (entry *Entry) ETracef(err error, format string, args ...interface{}) {
	entry.ELogf(TraceLevel, err, format, args...)
}
