import 'dart:io';
import 'package:flutter/services.dart';
import 'package:meta/meta.dart';

const MethodChannel _notifyChannel = MethodChannel('github.com/flutter_desktop_notify');

enum DesktopNotifyMode {
  notify,
  alert,
}

String _parseMode(DesktopNotifyMode mode) {
  switch (mode) {
    case DesktopNotifyMode.alert:
      return 'alert';
      break;
    case DesktopNotifyMode.notify:
      return 'notify';
      break;
    default:
      throw ArgumentError.value(mode);
  }
}

class DesktopNotify {
  static void show({
    @required String appName,
    @required String title,
    @required String text,
    String iconPath,
    DesktopNotifyMode mode = DesktopNotifyMode.notify,
  }) {
    if (Platform.isAndroid || Platform.isIOS) {
      // print('not supported platform: ${Platform.operatingSystem}');
      return;
    }

    _notifyChannel.invokeMethod(
      'notify',
      <String, dynamic>{
        'appName': appName,
        'title': title,
        'text': text,
        'iconPath': iconPath ?? '',
        'mode': _parseMode(mode),
      },
    );
  }
}
