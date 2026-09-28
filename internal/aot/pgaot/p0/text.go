package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_text_concat(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v5 = F_concat_internal(m, int32(_a_F_text_concat_0), int32(0), l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 == int32(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v5)
		}
	}
}
func F_text_format_append_string(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	if l3 == int32(0) {
		F_appendStringInfoString(m, l0, l1)
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			return
		}
	} else {
		if l3 < int32(0) {
			if l3 == int32(-2147483648) {
				F_errstart_cold(m, int32(21), int32(0))
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_errcode(m, int32(50331778))
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_text_format_append_string_0), int32(0))
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_text_format_append_string_1), int32(_a_F_text_format_append_string_2), int32(_a_F_text_format_append_string_3))
							v54 = m.ExcPending
							if v54 != 0 {
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
				v14 = F_pg_mbstrlen(m, l1)
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					v22 = int32(0) - l3
					v23 = v14
					F_appendStringInfoString(m, l0, l1)
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						if v22 <= v23 {
							return
						} else {
							F_appendStringInfoSpaces(m, l0, v22-v23)
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		} else {
			v16 = F_pg_mbstrlen(m, l1)
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				if l2&int32(1) == int32(0) {
					if l3 <= v16 {
						F_appendStringInfoString(m, l0, l1)
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							return
						}
					} else {
						F_appendStringInfoSpaces(m, l0, l3-v16)
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							F_appendStringInfoString(m, l0, l1)
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v22 = l3
					v23 = v16
					F_appendStringInfoString(m, l0, l1)
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						if v22 <= v23 {
							return
						} else {
							F_appendStringInfoSpaces(m, l0, v22-v23)
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_text_ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v17 == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
				if v23 == int32(18) {
					v26 = int32(16)
				} else {
					v26 = int32(0)
				}
				if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v33 = int32(4)
				} else {
					v33 = v26
				}
				v46 = v33
			} else {
				v34 = int32(1)
				if v17&v34 != 0 {
					v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v48 = int32(1)
			if v17&v48 != 0 {
				v52 = v48
			} else {
				v52 = int32(4)
			}
			v54 = int32(1)
			if v16&v54 != 0 {
				v58 = v54
			} else {
				v58 = int32(4)
			}
			if v16 == int32(1) {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v65 == int32(18) {
					v68 = int32(16)
				} else {
					v68 = int32(0)
				}
				if base.Ui32((v65-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v75 = int32(4)
				} else {
					v75 = v68
				}
				v88 = v75
			} else {
				v76 = int32(1)
				if v16&v76 != 0 {
					v88 = int32(base.Ui32(v16)>>(uint(v76)%32)) - v76
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v89 = F_varstr_cmp(m, v9+v52, v46, v14+v58, v88, v47)
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int64(0)
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v91 != v9 {
					F_pfree(m, v9)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int64(0)
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v95 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(int32(0) <= v89))
							}
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) <= v89))
						}
					}
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v95 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) <= v89))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(int32(0) <= v89))
					}
				}
			}
		}
	}
}
func F_text_pattern_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
			if v17 == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
				if v23 == int32(18) {
					v26 = int32(16)
				} else {
					v26 = int32(0)
				}
				if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v33 = int32(4)
				} else {
					v33 = v26
				}
				v46 = v33
			} else {
				v34 = int32(1)
				if v17&v34 != 0 {
					v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			if v47 == int32(1) {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
				if v53 == int32(18) {
					v56 = int32(16)
				} else {
					v56 = int32(0)
				}
				if base.Ui32((v53-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v63 = int32(4)
				} else {
					v63 = v56
				}
				v76 = v63
			} else {
				v64 = int32(1)
				if v47&v64 != 0 {
					v76 = int32(base.Ui32(v47)>>(uint(v64)%32)) - v64
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v77 = int32(1)
			if v17&v77 != 0 {
				v81 = v77
			} else {
				v81 = int32(4)
			}
			v83 = int32(1)
			if v47&v83 != 0 {
				v87 = v83
			} else {
				v87 = int32(4)
			}
			v89 = base.B2i32(v46 < v76)
			if v46 < v76 {
				v90 = v46
			} else {
				v90 = v76
			}
			v91 = F_memcmp(m, v6+v81, v11+v87, v90)
			mBase = m.M
			if v91 != 0 {
				v94 = v91
			} else {
				if v46 < v76 {
					v94 = int32(-1)
				} else {
					v94 = base.B2i32(v76 < v46)
				}
			}
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v95 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int64(0)
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v99 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) < v94))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(int32(0) < v94))
					}
				}
			} else {
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v99 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.B2i32(int32(0) < v94))
					}
				} else {
					return base.I64_extend_i32_u(base.B2i32(int32(0) < v94))
				}
			}
		}
	}
}
func F_text_regclass(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = F_textToQualifiedNameList(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int64(0)
		} else {
			v9 = F_makeRangeVarFromNameList(m, v7)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int64(0)
			} else {
				v11 = int32(0)
				v15 = F_RangeVarGetRelidExtended(m, v9, v11, v11, v11, v11)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v15)
				}
			}
		}
	}
}
func F_text_starts_with(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L5
	} else {
		goto L60
	}
L2:
	;
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_newlocale_from_collation(m, v8)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L55
	}
L5:
	;
	return int64(0)
L6:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v15 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v18 = F_toast_raw_datum_size(m, v10)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L9
	}
L8:
	;
	return base.I64_extend_i32_u(v144)
L9:
	;
	v20 = F_toast_raw_datum_size(m, v9)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if base.Ui32(v18) < base.Ui32(v20) {
		v144 = int32(0)
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v25 = F_text_substring(m, v10, int32(1), v20, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v28 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v9))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v30 = int32(1)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v32&v30 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v35 = v30
	goto L16
L15:
	;
	v35 = int32(4)
	goto L16
L16:
	;
	v36 = v25 + v35
	v37 = int32(1)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v41 = v39 & v37
	if v41 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v42 = v37
	goto L19
L18:
	;
	v42 = int32(4)
	goto L19
L19:
	;
	v43 = v28 + v42
	if v39 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v70) {
		goto L34
	} else {
		goto L35
	}
L21:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v49 == int32(18) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v60 = int32(1)
	if v41 != 0 {
		v70 = int32(base.Ui32(v39)>>(uint(v60)%32)) - v60
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v52 = int32(16)
	goto L26
L25:
	;
	v52 = int32(0)
	goto L26
L26:
	;
	if base.Ui32((v49-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v59 = int32(4)
	goto L29
L28:
	;
	v59 = v52
	goto L29
L29:
	;
	v70 = v59
	goto L20
L30:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v70 = int32(base.Ui32(v64)>>(uint(int32(2))%32)) - int32(4)
	goto L20
L31:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v133 != v25 {
		goto L49
	} else {
		goto L50
	}
L32:
	;
	v132 = int32(0)
	goto L31
L33:
	;
	v106 = v101
	v107 = v102
	v108 = v103
	goto L43
L34:
	;
	if (v36|v43)&int32(3) != 0 {
		v101 = v36
		v102 = v43
		v103 = v70
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v94 = v36
	v95 = v43
	v96 = v70
	goto L36
L36:
	;
	if v96 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L37:
	;
	v78 = v36
	v79 = v43
	v80 = v70
	goto L38
L38:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v83 != v84 {
		v101 = v78
		v102 = v79
		v103 = v80
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v94 = v89
	v95 = v87
	v96 = v91
	goto L36
L40:
	;
	v86 = int32(4)
	v87 = v79 + v86
	v89 = v78 + v86
	v91 = v80 - v86
	if base.Ui32(int32(3)) < base.Ui32(v91) {
		v78 = v89
		v79 = v87
		v80 = v91
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v101 = v94
	v102 = v95
	v103 = v96
	goto L33
L43:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v111 == v112 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v132 = v111 - v112
	goto L31
L45:
	;
	v114 = int32(1)
	v119 = v108 - v114
	if v119 != 0 {
		v106 = v106 + v114
		v107 = v107 + v114
		v108 = v119
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	goto L32
L49:
	;
	F_pfree(m, v25)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v138 = base.B2i32(v132 == int32(0))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v28 == v139 {
		v144 = v138
		goto L8
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	F_pfree(m, v28)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v144 = v138
	goto L8
L55:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_text_starts_with_0), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	F_errhint(m, int32(_a_F_text_starts_with_1), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_text_starts_with_2), int32(1337), int32(_a_F_text_starts_with_3))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(_a_F_text_starts_with_4), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_text_starts_with_2), int32(1610), int32(_a_F_text_starts_with_5))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_text_starts_with_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_like_regex_support(m, v2, int32(4))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_text_to_table(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(0)
	F_InitMaterializedSRF(m, l0, int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v18
		v22 = F_split_text(m, l0, v6+int32(4))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int64(0)
		} else {
			m.G0 = v6 + int32(16)
			return int64(0)
		}
	}
}
