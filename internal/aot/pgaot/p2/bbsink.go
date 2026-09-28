package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bbsink_copystream_begin_archive(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11+v12<<(uint(int32(2))%32))))
	F_pq_beginmessage(m, v7, int32(100))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		F_enlargeStringInfo(m, v7, int32(1))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v26 = int32(110)
			*(*uint8)(unsafe.Add(mBase, uint32(v23+v24))) = uint8(v26)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v23 + int32(1)
			F_pq_sendstring(m, v7, l1)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				if v33 != 0 {
					v35 = v33
				} else {
					v35 = int32(_a_F_bbsink_copystream_begin_archive_0)
				}
				F_pq_sendstring(m, v7, v35)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_pq_endmessage(m, v7)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_bbsink_copystream_begin_backup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int64
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int64
	_ = v174
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = F_palloc(m, v12+int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v15 + int32(7)
	v23 = int32(100)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+7)) = uint8(v23)
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	F_SendXlogRecPtrResult(m, v25, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v32 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_TupleDescInitBuiltinEntry(m, v35, int32(1), int32(_a_F_bbsink_copystream_begin_backup_0), int32(26))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitBuiltinEntry(m, v35, int32(2), int32(_a_F_bbsink_copystream_begin_backup_1), int32(25))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitBuiltinEntry(m, v35, int32(3), int32(_a_F_bbsink_copystream_begin_backup_2), int32(20))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v52 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v52 < v61 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v140 = F_begin_tup_output_tupdesc(m, v32, v35, int32(_a_F_bbsink_copystream_begin_backup_3))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L28
	}
L10:
	;
	v65 = v35 + int32(28)
	v72 = v52
	v73 = v61
	v75 = v52
	goto L14
L11:
	;
	v129 = v52
	v136 = v61
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v129
	goto L9
L13:
	;
	v129 = v123
	v136 = v102
	goto L12
L14:
	;
	v81 = v65 + v61<<(uint(int32(3))%32) + v72*int32(100)
	v84 = v65 + v72<<(uint(int32(3))%32)
	if v61 != v73 {
		v102 = v73
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v123 = v61
	goto L13
L16:
	;
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+2)))
	if v103 <= int32(0) {
		v123 = v72
		goto L13
	} else {
		goto L24
	}
L17:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+7)))
	if v86 != int32(118) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v102 = v72
	goto L16
L19:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+4)))
	if v89 != int32(1) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+6)))
	if v92&int32(6) != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v95 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+2)))
	if v95 <= int32(0) {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+90)))
	if v98 != int32(118) {
		v102 = v61
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L18
L24:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+90)))
	if v106 == int32(118) {
		v123 = v72
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+5)))
	v115 = (v75 + v109 - int32(1)) & (int32(0) - v109)
	if int32(_a_F_bbsink_copystream_begin_backup_4) < v115 {
		v123 = v72
		goto L13
	} else {
		goto L26
	}
L26:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v84))) = uint16(v115)
	v121 = v72 + int32(1)
	if v121 != v61 {
		v72 = v121
		v73 = v102
		v75 = v115 + v103
		goto L14
	} else {
		goto L27
	}
L27:
	;
	goto L15
L28:
	;
	if v29 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_end_tup_output(m, v140)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L45
	}
L30:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v144 <= int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v150 = int32(0)
	goto L32
L32:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153+v150<<(uint(int32(2))%32))))
	v158 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)) = uint8(v158)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)) = uint16(v158)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v162 == v158 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L29
L34:
	;
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v157)+16))
	if int64(0) <= v174 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v165 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)) = uint16(v165)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v167 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v157))))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v170 = F_cstring_to_text(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = base.I64_extend_i32_u(v170)
	goto L34
L39:
	;
	F_do_tup_output(m, v140, v9+int32(16), v9+int32(12))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(base.Ui64(v174) >> (uint(int64(10)) % 64))
	goto L39
L41:
	;
	goto L42
L42:
	;
	v180 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)) = uint8(v180)
	goto L39
L43:
	;
	v189 = v150 + int32(1)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v189 < v190 {
		v150 = v189
		goto L32
	} else {
		goto L44
	}
L44:
	;
	goto L33
L45:
	;
	F_pq_puttextmessage(m)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v203 = v9 + int32(16)
	F_pq_beginmessage(m, v203, int32(72))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_enlargeStringInfo(m, v203, int32(1))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v213 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v210+v211))) = uint8(v213)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v210 + int32(1)
	F_enlargeStringInfo(m, v203, int32(2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v224 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v221+v222))) = uint16(v224)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v221 + int32(2)
	F_pq_endmessage(m, v203)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	m.G0 = v9 + int32(48)
	return
}
func F_bbsink_copystream_manifest_contents(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v3 == int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_bbsink_copystream_manifest_contents[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		v13 = m.T0[v12].(func(*base.Module, int32, int32, int32) int32)(m, int32(100), v7, l1+int32(1))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_bbsink_forward_begin_archive(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	m.T0[v5].(func(*base.Module, int32, int32))(m, v3, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_bbsink_progress_begin_backup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int64
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v214 int32
	_ = v214
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_begin_backup[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v9
	v12 = *(*int64)(unsafe.Add(mBase, _c_F_bbsink_progress_begin_backup[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(3)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+24)))
	if v17 == int32(1) {
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
		v22 = v20
	} else {
		v22 = int64(-1)
	}
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v24 != 0 {
		v25 = int64(*(*int32)(unsafe.Add(mBase, uint32(v24)+4)))
		v27 = v25
	} else {
		v27 = int64(0)
	}
	*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v27
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_begin_backup[2]))
	if v41 == int32(0) {
	} else {
		v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_bbsink_progress_begin_backup[3])))
		if v45&int32(1) == int32(0) {
		} else {
			v50 = int32(_a_F_bbsink_progress_begin_backup_0)
			v52 = *(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_begin_backup[4]))
			v53 = int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_begin_backup[4])) = v52 + v53
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
			*(*int32)(unsafe.Add(mBase, uint32(v41))) = v56 + v53
			v60 = int32(0)
			v63 = base.AtomicRmwOr32(m, v60, int32(_a_F_bbsink_progress_begin_backup_1), v60)
			v155 = int32(0)
			v158 = int32(0)
			for {
				v167 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(32)+v158<<(uint(int32(2))%32))))
				v168 = int32(3)
				v174 = *(*int64)(unsafe.Add(mBase, uint32(v6+v158<<(uint(v168)%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v41+int32(232)+v167<<(uint(v168)%32)))) = v174
				v176 = int32(1)
				v179 = v155 + v176
				if v179 != int32(3) {
					v155 = v179
					v158 = v158 + v176
					continue
				} else {
					break
				}
				break
			}
			v190 = int32(0)
			v193 = base.AtomicRmwOr32(m, v190, int32(_a_F_bbsink_progress_begin_backup_1), v190)
			v194 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
			v195 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v41))) = v194 + v195
			v198 = int32(_a_F_bbsink_progress_begin_backup_0)
			v200 = *(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_begin_backup[4]))
			*(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_begin_backup[4])) = v200 - v195
		}
	}
	F_bbsink_forward_begin_backup(m, l0)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		return
	} else {
		m.G0 = v6 + int32(48)
		return
	}
}
func F_bbsink_progress_cleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_cleanup[0]))
	if v4 == int32(0) {
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_bbsink_progress_cleanup[1])))
		if v8&int32(1) == int32(0) {
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+220))
			if v13 == int32(0) {
			} else {
				v16 = int32(_a_F_bbsink_progress_cleanup_0)
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_cleanup[2]))
				v19 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_cleanup[2])) = v18 + v19
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v22 + v19
				v26 = int32(0)
				v28 = int32(_a_F_bbsink_progress_cleanup_1)
				v29 = base.AtomicRmwOr32(m, v26, v28, v26)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+220)) = v26
				*(*int32)(unsafe.Add(mBase, uint32(v4)+224)) = v26
				v37 = base.AtomicRmwOr32(m, v26, v28, v26)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v38 + v19
				v44 = *(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_cleanup[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_bbsink_progress_cleanup[2])) = v44 - v19
			}
		}
	}
	F_bbsink_forward_cleanup(m, l0)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		return
	} else {
		return
	}
}
func F_bbsink_server_manifest_contents(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v13
	v20 = F_FileWriteV(m, v12, v9+int32(40), int32(1), v11, int32(167772165))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		if l1 != v20 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				if v20 < int32(0) {
					F_errcode_for_file_access(m)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v71 = *(*int32)(unsafe.Add(mBase, _c_F_bbsink_server_manifest_contents[0]))
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v69*int32(48))+32))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v75
						F_errmsg(m, int32(_a_F_bbsink_server_manifest_contents_0), v9)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							F_errhint(m, int32(_a_F_bbsink_server_manifest_contents_1), int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_bbsink_server_manifest_contents_2), int32(268), int32(_a_F_bbsink_server_manifest_contents_3))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					F_errcode(m, int32(_a_F_bbsink_server_manifest_contents_4))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v34 = *(*int32)(unsafe.Add(mBase, _c_F_bbsink_server_manifest_contents[0]))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v32*int32(48))+32))
						v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v39
						*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v20
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v38
						F_errmsg(m, int32(_a_F_bbsink_server_manifest_contents_5), v9+int32(16))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							F_errhint(m, int32(_a_F_bbsink_server_manifest_contents_1), int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_bbsink_server_manifest_contents_2), int32(275), int32(_a_F_bbsink_server_manifest_contents_3))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			}
		} else {
			v58 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v58 + base.I64_extend_i32_s(l1)
			F_bbsink_forward_manifest_contents(m, l0, l1)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				m.G0 = v9 + int32(48)
				return
			}
		}
	}
}
