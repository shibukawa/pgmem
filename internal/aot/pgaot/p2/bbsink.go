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
					v35 = int32(735586)
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = F_palloc(m, v13+int32(8))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v16 + int32(8)
	v22 = v16 + int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v22
	v24 = int32(100)
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v24)
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	F_SendXlogRecPtrResult(m, v26, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v33 = F_CreateDestReceiver(m, int32(4))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_TupleDescInitBuiltinEntry(m, v36, int32(1), int32(428870), int32(26))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitBuiltinEntry(m, v36, int32(2), int32(262498), int32(25))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitBuiltinEntry(m, v36, int32(3), int32(337698), int32(20))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v54 = F_begin_tup_output_tupdesc(m, v33, v36, int32(1596068))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v30 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_end_tup_output(m, v54)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L27
	}
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v58 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v64 = int32(0)
	goto L13
L13:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v64<<(uint(int32(2))%32))))
	v73 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)) = uint8(v73)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v73)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v77 == v73 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L10
L15:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v72)+16))
	if int64(0) <= v87 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v80 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v80)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v82
	v84 = F_cstring_to_text(m, v77)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v84
	goto L15
L20:
	;
	F_do_tup_output(m, v54, v10+int32(16), v10+int32(12))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L25
	}
L21:
	;
	v92 = F_Int64GetDatum(m, int64(base.Ui64(v87)>>(uint(int64(10))%64)))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v95 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)) = uint8(v95)
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v92
	goto L20
L25:
	;
	v104 = v64 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v104 < v105 {
		v64 = v104
		goto L13
	} else {
		goto L26
	}
L26:
	;
	goto L14
L27:
	;
	F_pq_puttextmessage(m)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_pq_beginmessage(m, v10+int32(16), int32(72))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_enlargeStringInfo(m, v10+int32(16), int32(1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v131 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v128+v129))) = uint8(v131)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v128 + int32(1)
	F_enlargeStringInfo(m, v10+int32(16), int32(2))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v144 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v141+v142))) = uint16(v144)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v141 + int32(2)
	F_pq_endmessage(m, v10+int32(16))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	m.G0 = v10 + int32(32)
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
		v11 = *(*int32)(unsafe.Add(mBase, _consts[265]))
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int64
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v207 int32
	_ = v207
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[266]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v10
	v13 = *(*int64)(unsafe.Add(mBase, _consts[267]))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(3)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
	if v19 == int32(1) {
		v22 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
		v23 = v22
	} else {
		v23 = int64(-1)
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v25 != 0 {
		v26 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+4)))
		v28 = v26
	} else {
		v28 = int64(0)
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v28
	v33 = int32(0)
	v40 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	if v40 == v33 {
	} else {
		v46 = int32(*(*uint8)(unsafe.Add(mBase, _consts[24])))
		if v46&int32(1) == int32(0) {
		} else {
			v51 = int32(4470804)
			v53 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			v54 = int32(1)
			*(*int32)(unsafe.Add(mBase, _consts[7])) = v53 + v54
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
			*(*int32)(unsafe.Add(mBase, uint32(v40))) = v57 + v54
			v152 = int32(0)
			v155 = v33
			for {
				v164 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(32)+v155<<(uint(int32(2))%32))))
				v165 = int32(3)
				v171 = *(*int64)(unsafe.Add(mBase, uint32(v7+v155<<(uint(v165)%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v40+int32(232)+v164<<(uint(v165)%32)))) = v171
				v173 = int32(1)
				v176 = v152 + v173
				if v176 != int32(3) {
					v152 = v176
					v155 = v155 + v173
					continue
				} else {
					break
				}
				break
			}
			v187 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
			v188 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v40))) = v187 + v188
			v191 = int32(4470804)
			v193 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			*(*int32)(unsafe.Add(mBase, _consts[7])) = v193 - v188
		}
	}
	F_bbsink_forward_begin_backup(m, l0)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		return
	} else {
		m.G0 = v7 + int32(48)
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
						v71 = *(*int32)(unsafe.Add(mBase, _consts[262]))
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v69*int32(48))+32))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v75
						F_errmsg(m, int32(296090), v9)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							F_errhint(m, int32(622918), int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_errfinish(m, int32(489744), int32(268), int32(119274))
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
					F_errcode(m, int32(4293))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v34 = *(*int32)(unsafe.Add(mBase, _consts[262]))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v32*int32(48))+32))
						v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
						*(*uint32)(unsafe.Add(mBase, uint32(v9)+28)) = uint32(v39)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v20
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v38
						F_errmsg(m, int32(41019), v9+int32(16))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							F_errhint(m, int32(622918), int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								F_errfinish(m, int32(489744), int32(275), int32(119274))
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
