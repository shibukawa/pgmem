package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int8_avg_serialize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v164 int64
	_ = v164
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v173 int64
	_ = v173
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 == int32(0) {
		v40 = int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		switch v15 - int32(435) {
		case 0:
			v40 = int32(1)
		case 1:
			v40 = int32(2)
		default:
			v40 = int32(0)
		}
	}
	if v40 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_int8_avg_serialize_0), int32(0))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_int8_avg_serialize_1), int32(_a_F_int8_avg_serialize_2), int32(_a_F_int8_avg_serialize_3))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		F_pq_begintypsend(m, v8)
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int64(0)
		} else {
			v61 = *(*int64)(unsafe.Add(mBase, uint32(v58)+8))
			F_enlargeStringInfo(m, v8, int32(8))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int64(0)
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v68 = int64(56)
				v70 = int64(65280)
				v72 = int64(40)
				v75 = int64(16711680)
				v77 = int64(24)
				v79 = int64(4278190080)
				v81 = int64(8)
				*(*int64)(unsafe.Add(mBase, uint32(v65+v66))) = v61<<(uint(v68)%64) | v61&v70<<(uint(v72)%64) | (v61&v75<<(uint(v77)%64) | v61&v79<<(uint(v81)%64)) | (int64(base.Ui64(v61)>>(uint(v81)%64))&v79 | int64(base.Ui64(v61)>>(uint(v77)%64))&v75 | (int64(base.Ui64(v61)>>(uint(v72)%64))&v70 | int64(base.Ui64(v61)>>(uint(v68)%64))))
				v104 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v65 + v104
				v107 = *(*int64)(unsafe.Add(mBase, uint32(v58)+16))
				v108 = *(*int64)(unsafe.Add(mBase, uint32(v58)+24))
				F_enlargeStringInfo(m, v8, v104)
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int64(0)
				} else {
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v115 = int64(56)
					v117 = int64(65280)
					v119 = int64(40)
					v122 = int64(16711680)
					v124 = int64(24)
					v126 = int64(4278190080)
					v128 = int64(8)
					*(*int64)(unsafe.Add(mBase, uint32(v112+v113))) = v108<<(uint(v115)%64) | v108&v117<<(uint(v119)%64) | (v108&v122<<(uint(v124)%64) | v108&v126<<(uint(v128)%64)) | (int64(base.Ui64(v108)>>(uint(v128)%64))&v126 | int64(base.Ui64(v108)>>(uint(v124)%64))&v122 | (int64(base.Ui64(v108)>>(uint(v119)%64))&v117 | int64(base.Ui64(v108)>>(uint(v115)%64))))
					v151 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v112 + v151
					F_enlargeStringInfo(m, v8, v151)
					mBase = m.M
					v156 = m.ExcPending
					if v156 != 0 {
						return int64(0)
					} else {
						v157 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						v158 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v160 = int64(56)
						v162 = int64(65280)
						v164 = int64(40)
						v167 = int64(16711680)
						v169 = int64(24)
						v171 = int64(4278190080)
						v173 = int64(8)
						*(*int64)(unsafe.Add(mBase, uint32(v157+v158))) = v107<<(uint(v160)%64) | v107&v162<<(uint(v164)%64) | (v107&v167<<(uint(v169)%64) | v107&v171<<(uint(v173)%64)) | (int64(base.Ui64(v107)>>(uint(v173)%64))&v171 | int64(base.Ui64(v107)>>(uint(v169)%64))&v167 | (int64(base.Ui64(v107)>>(uint(v164)%64))&v162 | int64(base.Ui64(v107)>>(uint(v160)%64))))
						v197 = v157 + int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v197
						v200 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						*(*int32)(unsafe.Add(mBase, uint32(v200))) = v197 << (uint(int32(2)) % 32)
						m.G0 = v8 + int32(16)
						return base.I64_extend_i32_u(v200)
					}
				}
			}
		}
	}
}
func F_int8_decrement(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int64
	_ = v10
	v5 = base.B2i32(l1 == int64(-9223372036854775807-1))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v5)
	if l1 == int64(-9223372036854775807-1) {
		v10 = int64(0)
	} else {
		v10 = l1 - int64(1)
	}
	return v10
}
