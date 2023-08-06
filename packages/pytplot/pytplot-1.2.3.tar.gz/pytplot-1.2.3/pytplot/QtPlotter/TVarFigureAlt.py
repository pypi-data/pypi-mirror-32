import pyqtgraph as pg
import numpy as np
from .. import tplot_utilities 
from pytplot import tplot_opt_glob
import pytplot
from pyqtgraph.Qt import QtCore
from .CustomAxis.BlankAxis import BlankAxis


class TVarFigureAlt(pg.GraphicsLayout):
    def __init__(self, tvar_name, show_xaxis=False, mouse_function=None):
        
        self.tvar_name=tvar_name
        self.show_xaxis = show_xaxis
        
        #Sets up the layout of the Tplot Object
        pg.GraphicsLayout.__init__(self)
        self.layout.setHorizontalSpacing(50)
        self.layout.setContentsMargins(0,0,0,0)
        #Set up the x axis
        self.xaxis = pg.AxisItem(orientation='bottom')
        self.xaxis.setHeight(35)
        self.xaxis.enableAutoSIPrefix(enable=False)
        #Set up the y axis
        self.yaxis = pg.AxisItem("left")
        self.yaxis.setWidth(100)
        
        self.plotwindow = self.addPlot(row=0, col=0, axisItems={'bottom': self.xaxis, 'left': self.yaxis})
        
        #Set up the view box needed for the legends
        self.legendvb = pg.ViewBox(enableMouse=False)
        self.legendvb.setMaximumWidth(100)
        self.legendvb.setXRange(0,1, padding=0)
        self.legendvb.setYRange(0,1, padding=0)
        self.addItem(self.legendvb,0,1)       
        
        
        self.curves = []
        self.colors = self._setcolors()
        self.colormap = self._setcolormap()

        if show_xaxis:
            self.plotwindow.showAxis('bottom')
        else:
            self.plotwindow.hideAxis('bottom')
        
        self._mouseMovedFunction = mouse_function
    
    
    def buildfigure(self):
        self._setxrange()
        self._setyrange()
        self._setyaxistype()
        self._setzaxistype()
        self._setzrange()
        self._visdata()
        self._addtimebars()
        self._setyaxislabel()
        self._setxaxislabel()
        self._addmouseevents()
        self._addlegend()
    
    def getfig(self):
        return self
    
    def getaxistype(self):
        axis_type = 'alt'
        link_y_axis = False
        return axis_type, link_y_axis
    
    def _setyaxislabel(self):
        self.yaxis.setLabel(pytplot.data_quants[self.tvar_name].yaxis_opt['axis_label'])
    
    def _setyaxistype(self):
        if self._getyaxistype() == 'log':
            self.plotwindow.setLogMode(y=True)
        else:
            self.plotwindow.setLogMode(y=False)
        return
    
    def _getyaxistype(self):
        if 'y_axis_type' in pytplot.data_quants[self.tvar_name].yaxis_opt:
            return pytplot.data_quants[self.tvar_name].yaxis_opt['y_axis_type']
        else:
            return 'linear'
    
    def _setxrange(self):
        if 'alt_range' in tplot_opt_glob:
            self.plotwindow.setXRange(tplot_opt_glob['alt_range'][0], tplot_opt_glob['alt_range'][1])
        else:
            return
    
    def _setxaxislabel(self):
        self.xaxis.setLabel("Altitude")
        
    def getaxistype(self):
        axis_type = 'altitude'
        link_y_axis = False
        return axis_type, link_y_axis
    
    def _addmouseevents(self):
        return
    
    def _addlegend(self):
        if 'legend_names' in pytplot.data_quants[self.tvar_name].yaxis_opt:
            legend_names = pytplot.data_quants[self.tvar_name].yaxis_opt['legend_names']
            if len(legend_names) != len(self.curves):
                print("Number of lines do not match length of legend names")
            if len(legend_names) == 1:
                pos_array=[.5]
            else:
                pos_array=np.linspace(1,0,len(legend_names))
            i=0
            for legend_name in legend_names:
                if i+1 == len(legend_names): #Last
                    text = pg.TextItem(text=legend_name, anchor=(0,1.5), color=self.colors[i])
                elif i==0: #First
                    text = pg.TextItem(text=legend_name, anchor=(0,-.5), color=self.colors[i])
                else: #All others
                    text = pg.TextItem(text=legend_name, anchor=(0,0.5), color=self.colors[i])
                self.legendvb.addItem(text)
                text.setPos(0,pos_array[i])
                i+=1
                
    def _setzaxistype(self):
        if self._getzaxistype() == 'log':
            self.zscale = 'log'
        else:
            self.zscale = 'linear'
    
    def _getzaxistype(self):
        return
            
    def _setcolors(self):
        if 'line_color' in pytplot.data_quants[self.tvar_name].extras:
            return pytplot.data_quants[self.tvar_name].extras['line_color']
        else: 
            return ['k', 'r', 'g', 'c', 'y', 'm', 'b']
    
    def _setcolormap(self):          
        return
    
    def _setyrange(self):
        if self._getyaxistype() == 'log':
            if pytplot.data_quants[self.tvar_name].yaxis_opt['y_range'][0] <0 or pytplot.data_quants[self.tvar_name].yaxis_opt['y_range'][1] < 0:
                return
            self.plotwindow.vb.setYRange(np.log10(pytplot.data_quants[self.tvar_name].yaxis_opt['y_range'][0]), np.log10(pytplot.data_quants[self.tvar_name].yaxis_opt['y_range'][1]), padding=0)
        else:
            self.plotwindow.vb.setYRange(pytplot.data_quants[self.tvar_name].yaxis_opt['y_range'][0], pytplot.data_quants[self.tvar_name].yaxis_opt['y_range'][1], padding=0)
    
    def _setzrange(self):
        return
    
    def _addtimebars(self):
        #initialize dataset variable
        datasets = []
        #grab tbardict
        tbardict = pytplot.data_quants[self.tvar_name].time_bar
        ltbar = len(tbardict)
        #make sure data is in list format
        if isinstance(pytplot.data_quants[self.tvar_name].data, list):
            for oplot_name in pytplot.data_quants[self.tvar_name].data:
                datasets.append(pytplot.data_quants[oplot_name])
        else:
            datasets.append(pytplot.data_quants[self.tvar_name])        
        for dataset in datasets:  
            #for location in tbar dict
            for i in range(ltbar):
                #get times, color, point size
                test_time = pytplot.data_quants[self.tvar_name].time_bar[i]["location"]
                print(test_time)
                color = pytplot.data_quants[self.tvar_name].time_bar[i]["line_color"]
                pointsize = pytplot.data_quants[self.tvar_name].time_bar[i]["line_width"]
                #correlate given time with corresponding data/alt points
                time, altitude = pytplot.get_data(dataset.links['alt']) 
                altitude = altitude.transpose()[0]
                nearest_time_index = np.abs(time - test_time).argmin()
                data_point = dataset.data.iloc[nearest_time_index][0]
                alt_point = altitude[nearest_time_index]
                print(alt_point,data_point)
                self.plotwindow.scatterPlot([alt_point], [data_point], size = pointsize, pen=pg.mkPen(None), brush=color)
        return
    def _visdata(self):
        datasets = []
        if isinstance(pytplot.data_quants[self.tvar_name].data, list):
            for oplot_name in pytplot.data_quants[self.tvar_name].data:
                datasets.append(pytplot.data_quants[oplot_name])
        else:
            datasets.append(pytplot.data_quants[self.tvar_name])
        line_num = 0
        for dataset in datasets:  
            for i in range(0,len(dataset.data.columns)):
                _, x = pytplot.get_data(dataset.links['alt']) 
                x = x.transpose()[0]
                self.curves.append(self.plotwindow.scatterPlot(x.tolist(), dataset.data[i].tolist(), 
                                                        pen=pg.mkPen(None), brush=self.colors[line_num % len(self.colors)]))
                line_num+=1