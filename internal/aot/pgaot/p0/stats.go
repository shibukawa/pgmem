package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_StatsShmemSize(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	v5 = F_add_size(m, int32(_a_F_StatsShmemSize_0), int32(_a_F_StatsShmemSize_1))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v24 = int32(0)
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemSize[0]))
		if v26 == v24 {
			v36 = v24
		} else {
			v30 = int32(96)
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v26+v30-v30)))
			v36 = v34
		}
		if v36 == int32(0) {
			v52 = v5
		} else {
			v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
			if v41&int32(1) == int32(0) {
				v52 = v5
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
				v52 = (v46+int32(7))&int32(-8) + v5
			}
		}
		v68 = int32(0)
		v70 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemSize[0]))
		if v70 == v68 {
			v80 = v68
		} else {
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v70+int32(100)-int32(96))))
			v80 = v78
		}
		if v80 == int32(0) {
			v96 = v52
		} else {
			v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
			if v85&int32(1) == int32(0) {
				v96 = v52
			} else {
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
				v96 = (v90+int32(7))&int32(-8) + v52
			}
		}
		v112 = int32(0)
		v114 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemSize[0]))
		if v114 == v112 {
			v124 = v112
		} else {
			v122 = *(*int32)(unsafe.Add(mBase, uint32(v114+int32(104)-int32(96))))
			v124 = v122
		}
		if v124 == int32(0) {
			v140 = v96
		} else {
			v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
			if v129&int32(1) == int32(0) {
				v140 = v96
			} else {
				v134 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
				v140 = (v134+int32(7))&int32(-8) + v96
			}
		}
		v156 = int32(0)
		v158 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemSize[0]))
		if v158 == v156 {
			v168 = v156
		} else {
			v166 = *(*int32)(unsafe.Add(mBase, uint32(v158+int32(108)-int32(96))))
			v168 = v166
		}
		if v168 == int32(0) {
			v184 = v140
		} else {
			v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
			if v173&int32(1) == int32(0) {
				v184 = v140
			} else {
				v178 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
				v184 = (v178+int32(7))&int32(-8) + v140
			}
		}
		v200 = int32(0)
		v202 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemSize[0]))
		if v202 == v200 {
			v212 = v200
		} else {
			v210 = *(*int32)(unsafe.Add(mBase, uint32(v202+int32(112)-int32(96))))
			v212 = v210
		}
		if v212 == int32(0) {
			v228 = v184
		} else {
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
			if v217&int32(1) == int32(0) {
				v228 = v184
			} else {
				v222 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
				v228 = (v222+int32(7))&int32(-8) + v184
			}
		}
		v244 = int32(0)
		v246 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemSize[0]))
		if v246 == v244 {
			v256 = v244
		} else {
			v254 = *(*int32)(unsafe.Add(mBase, uint32(v246+int32(116)-int32(96))))
			v256 = v254
		}
		if v256 == int32(0) {
			v272 = v228
		} else {
			v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256))))
			if v261&int32(1) == int32(0) {
				v272 = v228
			} else {
				v266 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
				v272 = (v266+int32(7))&int32(-8) + v228
			}
		}
		v288 = int32(0)
		v290 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemSize[0]))
		if v290 == v288 {
			v300 = v288
		} else {
			v298 = *(*int32)(unsafe.Add(mBase, uint32(v290+int32(120)-int32(96))))
			v300 = v298
		}
		if v300 == int32(0) {
			v316 = v272
		} else {
			v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
			if v305&int32(1) == int32(0) {
				v316 = v272
			} else {
				v310 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
				v316 = (v310+int32(7))&int32(-8) + v272
			}
		}
		v332 = int32(0)
		v334 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemSize[0]))
		if v334 == v332 {
			v344 = v332
		} else {
			v342 = *(*int32)(unsafe.Add(mBase, uint32(v334+int32(124)-int32(96))))
			v344 = v342
		}
		if v344 == int32(0) {
			v360 = v316
		} else {
			v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
			if v349&int32(1) == int32(0) {
				v360 = v316
			} else {
				v354 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
				v360 = (v354+int32(7))&int32(-8) + v316
			}
		}
		v376 = int32(0)
		v378 = *(*int32)(unsafe.Add(mBase, _c_F_StatsShmemSize[0]))
		if v378 == v376 {
			v388 = v376
		} else {
			v386 = *(*int32)(unsafe.Add(mBase, uint32(v378+int32(128)-int32(96))))
			v388 = v386
		}
		if v388 == int32(0) {
			v404 = v360
		} else {
			v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
			if v393&int32(1) == int32(0) {
				v404 = v360
			} else {
				v398 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
				v404 = (v398+int32(7))&int32(-8) + v360
			}
		}
		return v404
	}
}
