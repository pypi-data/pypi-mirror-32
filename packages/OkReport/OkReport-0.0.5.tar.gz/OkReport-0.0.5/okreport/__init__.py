# !/usr/bin/python -u

"""
Copyright (C) 2018 LingoChamp Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""
import argparse
from os import listdir, remove, makedirs

from os.path import exists

from shutil import rmtree

from okreport import html, sender
from okreport.config import Config
from okreport.helper import handle_home_case
from okreport.report import Report

__author__ = 'JacksGong'
__version__ = '0.0.5'
__description__ = 'This tool is used for assembling report from okcheck and send to mails'


def main():
    print("-------------------------------------------------------")
    print("                  OkReport v" + __version__)
    print("")
    print(__description__)
    print("")
    print("                   Have Fun!")
    print("-------------------------------------------------------")

    parser = argparse.ArgumentParser(description=__description__)
    parser.add_argument('report_path', nargs='*',
                        help='This can be Application package name(s) or log file path(if the file from path is exist)',
                        default='./build/reports')
    parser.add_argument('-c', '--config_path', dest='config_path', help='config file path',
                        default='~/.okreport/.config')
    parser.add_argument('-r', '--report_cache_path', dest='report_cache_path', help='report cache file path',
                        default='~/.okreport/')
    parser.add_argument('-d', '--dump', dest='dump', action='store_true', help='just dump no send')

    args = parser.parse_args()

    report_path = handle_home_case(args.report_path[0])
    report_cache_path = handle_home_case(args.report_cache_path)
    config_path = handle_home_case(args.config_path)
    if not exists(config_path):
        print "config path %s isn't exist!" % config_path
        exit(-1)

    _report = Report()
    _report.fromConfig(report_cache_path)
    _report.addParse(report_path)

    buildDir = "build"
    chartDir = "%s/okreport" % buildDir
    needRemoveBuildDir = not exists(buildDir)

    mailConfig = Config(config_path)

    if not exists(chartDir):
        makedirs(chartDir)

    try:
        chartPathList = list()
        reportHtml = html.BASIC_REPORT_HTML
        reportHtml += '<h2>OkPmd</h2>\n'
        reportHtml += _report.getPmdChartHtmlFragment(chartPathList, chartDir)
        reportHtml += '<h2>OkFindbugs</h2>\n'
        reportHtml += _report.getFindbugsChartHtmlFragment(chartPathList, chartDir)
        reportHtml += '<h2>OkLint</h2>\n'
        reportHtml += _report.getLintChartHtmlFragment(chartPathList, chartDir)

        if mailConfig.contentUrl():
            reportHtml += '<p>For more detail please move to: <a href="%s">Here</a></p>\n' % mailConfig.contentUrl()

        reportHtml += html.END_HTML

        _report.addLastConfig(report_cache_path)
        _report.maintainConfigs(report_cache_path)

        _report.dump()

        if not args.dump:
            sender.send_mail(reportHtml, chartPathList, mailConfig)
    finally:
        if needRemoveBuildDir:
            print "remove build cache dir: %s" % buildDir
            rmtree(buildDir)
        else:
            print "remove chart cache dir: %s" % chartDir
            rmtree(chartDir)
