package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_expanded_record_fetch_tupdesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v4 == int32(0) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v9 = F_lookup_rowtype_tupdesc(m, v7, v8)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			if int32(0) <= v13 {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
				if v16 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1292)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = l0
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v24 = l0 + int32(108)
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v25
					v27 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)) = uint8(v27)
					*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v24
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v9
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				v33 = v31 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v33
				if v33 < int32(0) {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					v43 = F_assign_record_type_identifier(m, v41, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v43
						v46 = v9
						return v46
					}
				} else {
					F_DecrTupleDescRefCount(m, v9)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						v43 = F_assign_record_type_identifier(m, v41, v42)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v43
							v46 = v9
							return v46
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v9
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				v43 = F_assign_record_type_identifier(m, v41, v42)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v43
					v46 = v9
					return v46
				}
			}
		}
	} else {
		v46 = v4
		return v46
	}
}
func F_expanded_record_lookup_field(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = F_expanded_record_fetch_tupdesc(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v13 = v6
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if int32(0) < v14 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	v13 = v9
	goto L3
L6:
	;
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v99)+74)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v99)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99)+76))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v99)+96))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v107
	return int32(1)
L7:
	;
	v18 = int32(0)
	v21 = v14
	goto L10
L8:
	;
	goto L9
L9:
	;
	v64 = F_strcmp(m, int32(_a_F_expanded_record_lookup_field_0), l1)
	mBase = m.M
	if v64 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L10:
	;
	v28 = v13 + v21<<(uint(int32(4))%32) + v18*int32(100)
	v30 = v28 + int32(24)
	if v30|l1 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L9
L12:
	;
	if v45 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L13:
	;
	v36 = int32(-1)
	goto L15
L14:
	;
	v36 = int32(0)
	goto L15
L15:
	;
	if v30 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v37 = int32(1)
	goto L18
L17:
	;
	v37 = v36
	goto L18
L18:
	;
	v38 = int32(0)
	if base.B2i32(v30 == v38)|base.B2i32(l1 == v38) != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v45 = v37
	goto L21
L20:
	;
	v44 = F_strncmp(m, v30, l1, int32(64))
	mBase = m.M
	v45 = v44
	goto L21
L21:
	;
	goto L12
L22:
	;
	v49 = v28 + int32(20)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+91)))
	if v50 != int32(1) {
		v99 = v49
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v55 = v18 + int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v55 < v56 {
		v18 = v55
		v21 = v56
		goto L10
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	goto L11
L27:
	;
	if v93 != 0 {
		v99 = v93
		goto L6
	} else {
		goto L46
	}
L28:
	;
	v93 = int32(_a_F_expanded_record_lookup_field_1)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v69 = F_strcmp(m, int32(_a_F_expanded_record_lookup_field_2), l1)
	mBase = m.M
	if v69 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v93 = int32(_a_F_expanded_record_lookup_field_3)
	goto L27
L32:
	;
	goto L33
L33:
	;
	v74 = F_strcmp(m, int32(_a_F_expanded_record_lookup_field_4), l1)
	mBase = m.M
	if v74 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v93 = int32(_a_F_expanded_record_lookup_field_5)
	goto L27
L35:
	;
	goto L36
L36:
	;
	v79 = F_strcmp(m, int32(_a_F_expanded_record_lookup_field_6), l1)
	mBase = m.M
	if v79 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v93 = int32(_a_F_expanded_record_lookup_field_7)
	goto L27
L38:
	;
	goto L39
L39:
	;
	v84 = F_strcmp(m, int32(_a_F_expanded_record_lookup_field_8), l1)
	mBase = m.M
	if v84 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v93 = int32(_a_F_expanded_record_lookup_field_9)
	goto L27
L41:
	;
	goto L42
L42:
	;
	v91 = F_strcmp(m, int32(_a_F_expanded_record_lookup_field_10), l1)
	mBase = m.M
	if v91 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v92 = int32(0)
	goto L45
L44:
	;
	v92 = int32(_a_F_expanded_record_lookup_field_11)
	goto L45
L45:
	;
	v93 = v92
	goto L27
L46:
	;
	return int32(0)
}
func F_make_expanded_record_from_typeid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 != int32(2249) {
		v15 = F_lookup_type_cache(m, l0, int32(_a_F_make_expanded_record_from_typeid_0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+13)))
			if v19 == int32(100) {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
				v25 = F_lookup_type_cache(m, v23, int32(256))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = v25
					v28 = int32(64)
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+188))
					if v29 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(151027844))
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return int32(0)
							} else {
								v128 = F_format_type_be(m, l0)
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v128
									F_errmsg(m, int32(_a_F_make_expanded_record_from_typeid_1), v10)
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_make_expanded_record_from_typeid_2), int32(100), int32(_a_F_make_expanded_record_from_typeid_3))
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					} else {
						v32 = *(*int64)(unsafe.Add(mBase, uint32(v27)+192))
						v40 = v29
						v41 = v28
						v42 = v32
						v47 = F_AllocSetContextCreateInternal(m, l2, int32(_a_F_make_expanded_record_from_typeid_4), int32(0), int32(_a_F_make_expanded_record_from_typeid_5), int32(_a_F_make_expanded_record_from_typeid_6))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
							v54 = F_MemoryContextAlloc(m, v47, v49*int32(5)+int32(120))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								base.MemoryFill(m, v54, int32(0), int32(120))
								v60 = int32(513)
								*(*uint16)(unsafe.Add(mBase, uint32(v54)+18)) = uint16(v60)
								v62 = int32(769)
								*(*uint16)(unsafe.Add(mBase, uint32(v54)+12)) = uint16(v62)
								*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v47
								*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = int32(_a_F_make_expanded_record_from_typeid_7)
								*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v54
								*(*int32)(unsafe.Add(mBase, uint32(v54)+14)) = v54
								v71 = v54 + int32(120)
								*(*int32)(unsafe.Add(mBase, uint32(v54)+56)) = v71
								*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = int32(1384727874)
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
								*(*int32)(unsafe.Add(mBase, uint32(v54)+60)) = v71 + v75<<(uint(int32(2))%32)
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
								*(*int32)(unsafe.Add(mBase, uint32(v54)+32)) = l0
								*(*int32)(unsafe.Add(mBase, uint32(v54)+64)) = v80
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v54)+36)) = v83
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v54)+48)) = v42
								*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v85
								*(*int32)(unsafe.Add(mBase, uint32(v54)+28)) = v41
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
								if int32(0) <= v89 {
									*(*int32)(unsafe.Add(mBase, uint32(v54)+108)) = int32(1292)
									*(*int32)(unsafe.Add(mBase, uint32(v54)+112)) = v54
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
									v97 = v54 + int32(108)
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v98
									v100 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v95)+4)) = uint8(v100)
									*(*int32)(unsafe.Add(mBase, uint32(v95)+40)) = v97
									*(*int32)(unsafe.Add(mBase, uint32(v54)+44)) = v40
									v104 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
									v106 = v104 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v106
									if base.B2i32(l0 != int32(2249))|base.B2i32(v106 < int32(0)) != 0 {
										m.G0 = v10 + int32(16)
										return v54
									} else {
										F_DecrTupleDescRefCount(m, v40)
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											m.G0 = v10 + int32(16)
											return v54
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v54)+44)) = v40
									m.G0 = v10 + int32(16)
									return v54
								}
							}
						}
					}
				}
			} else {
				v27 = v15
				v28 = v4
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+188))
				if v29 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return int32(0)
						} else {
							v128 = F_format_type_be(m, l0)
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v128
								F_errmsg(m, int32(_a_F_make_expanded_record_from_typeid_1), v10)
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_make_expanded_record_from_typeid_2), int32(100), int32(_a_F_make_expanded_record_from_typeid_3))
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					v32 = *(*int64)(unsafe.Add(mBase, uint32(v27)+192))
					v40 = v29
					v41 = v28
					v42 = v32
					v47 = F_AllocSetContextCreateInternal(m, l2, int32(_a_F_make_expanded_record_from_typeid_4), int32(0), int32(_a_F_make_expanded_record_from_typeid_5), int32(_a_F_make_expanded_record_from_typeid_6))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						v54 = F_MemoryContextAlloc(m, v47, v49*int32(5)+int32(120))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							base.MemoryFill(m, v54, int32(0), int32(120))
							v60 = int32(513)
							*(*uint16)(unsafe.Add(mBase, uint32(v54)+18)) = uint16(v60)
							v62 = int32(769)
							*(*uint16)(unsafe.Add(mBase, uint32(v54)+12)) = uint16(v62)
							*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v47
							*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = int32(_a_F_make_expanded_record_from_typeid_7)
							*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v54
							*(*int32)(unsafe.Add(mBase, uint32(v54)+14)) = v54
							v71 = v54 + int32(120)
							*(*int32)(unsafe.Add(mBase, uint32(v54)+56)) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = int32(1384727874)
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
							*(*int32)(unsafe.Add(mBase, uint32(v54)+60)) = v71 + v75<<(uint(int32(2))%32)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
							*(*int32)(unsafe.Add(mBase, uint32(v54)+32)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v54)+64)) = v80
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v54)+36)) = v83
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v54)+48)) = v42
							*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v85
							*(*int32)(unsafe.Add(mBase, uint32(v54)+28)) = v41
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
							if int32(0) <= v89 {
								*(*int32)(unsafe.Add(mBase, uint32(v54)+108)) = int32(1292)
								*(*int32)(unsafe.Add(mBase, uint32(v54)+112)) = v54
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
								v97 = v54 + int32(108)
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v98
								v100 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v95)+4)) = uint8(v100)
								*(*int32)(unsafe.Add(mBase, uint32(v95)+40)) = v97
								*(*int32)(unsafe.Add(mBase, uint32(v54)+44)) = v40
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
								v106 = v104 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v106
								if base.B2i32(l0 != int32(2249))|base.B2i32(v106 < int32(0)) != 0 {
									m.G0 = v10 + int32(16)
									return v54
								} else {
									F_DecrTupleDescRefCount(m, v40)
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return v54
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v54)+44)) = v40
								m.G0 = v10 + int32(16)
								return v54
							}
						}
					}
				}
			}
		}
	} else {
		v34 = F_lookup_rowtype_tupdesc(m, int32(2249), l1)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			v37 = F_assign_record_type_identifier(m, int32(2249), l1)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v40 = v34
				v41 = v4
				v42 = v37
				v47 = F_AllocSetContextCreateInternal(m, l2, int32(_a_F_make_expanded_record_from_typeid_4), int32(0), int32(_a_F_make_expanded_record_from_typeid_5), int32(_a_F_make_expanded_record_from_typeid_6))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					v54 = F_MemoryContextAlloc(m, v47, v49*int32(5)+int32(120))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						base.MemoryFill(m, v54, int32(0), int32(120))
						v60 = int32(513)
						*(*uint16)(unsafe.Add(mBase, uint32(v54)+18)) = uint16(v60)
						v62 = int32(769)
						*(*uint16)(unsafe.Add(mBase, uint32(v54)+12)) = uint16(v62)
						*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v47
						*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = int32(_a_F_make_expanded_record_from_typeid_7)
						*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v54
						*(*int32)(unsafe.Add(mBase, uint32(v54)+14)) = v54
						v71 = v54 + int32(120)
						*(*int32)(unsafe.Add(mBase, uint32(v54)+56)) = v71
						*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = int32(1384727874)
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						*(*int32)(unsafe.Add(mBase, uint32(v54)+60)) = v71 + v75<<(uint(int32(2))%32)
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						*(*int32)(unsafe.Add(mBase, uint32(v54)+32)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v54)+64)) = v80
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v54)+36)) = v83
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v54)+48)) = v42
						*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v85
						*(*int32)(unsafe.Add(mBase, uint32(v54)+28)) = v41
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
						if int32(0) <= v89 {
							*(*int32)(unsafe.Add(mBase, uint32(v54)+108)) = int32(1292)
							*(*int32)(unsafe.Add(mBase, uint32(v54)+112)) = v54
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
							v97 = v54 + int32(108)
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v98
							v100 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v95)+4)) = uint8(v100)
							*(*int32)(unsafe.Add(mBase, uint32(v95)+40)) = v97
							*(*int32)(unsafe.Add(mBase, uint32(v54)+44)) = v40
							v104 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
							v106 = v104 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v106
							if base.B2i32(l0 != int32(2249))|base.B2i32(v106 < int32(0)) != 0 {
								m.G0 = v10 + int32(16)
								return v54
							} else {
								F_DecrTupleDescRefCount(m, v40)
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(16)
									return v54
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v54)+44)) = v40
							m.G0 = v10 + int32(16)
							return v54
						}
					}
				}
			}
		}
	}
}
