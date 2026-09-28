package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AsyncExistsPendingNotify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l0
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncExistsPendingNotify[0]))
	if v18 == v2 {
		v145 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v145
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = int32(0)
	v26 = F_hash_search(m, v21, v14+int32(12), v24, v24)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v33 == int32(0) {
		v145 = v2
		goto L1
	} else {
		goto L9
	}
L6:
	;
	return int32(0)
L7:
	;
	if v26 == int32(0) {
		v145 = v2
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v145 = int32(1)
	goto L1
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v36 <= int32(0) {
		v145 = v2
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v39 = int32(0)
	if v39 < v36 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v43 = v36
	goto L13
L12:
	;
	v43 = v39
	goto L13
L13:
	;
	v45 = l0 + int32(4)
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v52 = v39
	goto L14
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v49+v52<<(uint(int32(2))%32))))
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64))))
	if v46 != v65 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v145 = v2
	goto L1
L16:
	;
	v138 = v52 + int32(1)
	if v138 != v43 {
		v52 = v138
		goto L14
	} else {
		goto L38
	}
L17:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+2)))
	if v67 != v68 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v70 = int32(4)
	v71 = v64 + v70
	v72 = v46 + int32(2) + v67
	if base.Ui32(v70) <= base.Ui32(v72) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	if v134 != 0 {
		goto L16
	} else {
		goto L37
	}
L20:
	;
	v134 = int32(0)
	goto L19
L21:
	;
	v108 = v103
	v109 = v104
	v110 = v105
	goto L31
L22:
	;
	if (v45|v71)&int32(3) != 0 {
		v103 = v45
		v104 = v71
		v105 = v72
		goto L21
	} else {
		goto L25
	}
L23:
	;
	v96 = v45
	v97 = v71
	v98 = v72
	goto L24
L24:
	;
	if v98 == int32(0) {
		goto L20
	} else {
		goto L30
	}
L25:
	;
	v80 = v45
	v81 = v71
	v82 = v72
	goto L26
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v85 != v86 {
		v103 = v80
		v104 = v81
		v105 = v82
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v96 = v91
	v97 = v89
	v98 = v93
	goto L24
L28:
	;
	v88 = int32(4)
	v89 = v81 + v88
	v91 = v80 + v88
	v93 = v82 - v88
	if base.Ui32(int32(3)) < base.Ui32(v93) {
		v80 = v91
		v81 = v89
		v82 = v93
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v103 = v96
	v104 = v97
	v105 = v98
	goto L21
L31:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v113 == v114 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v134 = v113 - v114
	goto L19
L33:
	;
	v116 = int32(1)
	v121 = v110 - v116
	if v121 != 0 {
		v108 = v108 + v116
		v109 = v109 + v116
		v110 = v121
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L20
L37:
	;
	v145 = int32(1)
	goto L1
L38:
	;
	goto L15
}
func F_ExecAsyncRequest(m *base.Module, l0 int32) {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
	if v10 != 0 {
		F_ExecReScan(m, v9)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = v13
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
			if v15 != 0 {
				F_InstrStart(m, v15)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v19 = v18
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					if v20 == int32(424) {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+128))
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+176))
						m.T0[v25].(func(*base.Module, int32))(m, l0)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
							if v29 == int32(403) {
								F_ExecAsyncAppendResponse(m, l0)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
									if v35 != 0 {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v36 != 0 {
											v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
											if v39&int32(2) != 0 {
												v42 = float64(0)
											} else {
												v42 = float64(1)
											}
											v44 = v42
										} else {
											v44 = float64(0)
										}
										F_InstrStopNode(m, v35, v44)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											m.G0 = v7 + int32(32)
											return
										}
									} else {
										m.G0 = v7 + int32(32)
										return
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v56
									F_errmsg_internal(m, int32(_a_F_ExecAsyncRequest_0), v7)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_ExecAsyncRequest_1), int32(128), int32(_a_F_ExecAsyncRequest_2))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
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
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v71
							F_errmsg_internal(m, int32(_a_F_ExecAsyncRequest_0), v7+int32(16))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ExecAsyncRequest_1), int32(44), int32(_a_F_ExecAsyncRequest_3))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
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
			} else {
				v19 = v14
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				if v20 == int32(424) {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+128))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+176))
					m.T0[v25].(func(*base.Module, int32))(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						if v29 == int32(403) {
							F_ExecAsyncAppendResponse(m, l0)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
								if v35 != 0 {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v36 != 0 {
										v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
										if v39&int32(2) != 0 {
											v42 = float64(0)
										} else {
											v42 = float64(1)
										}
										v44 = v42
									} else {
										v44 = float64(0)
									}
									F_InstrStopNode(m, v35, v44)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										m.G0 = v7 + int32(32)
										return
									}
								} else {
									m.G0 = v7 + int32(32)
									return
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v56
								F_errmsg_internal(m, int32(_a_F_ExecAsyncRequest_0), v7)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecAsyncRequest_1), int32(128), int32(_a_F_ExecAsyncRequest_2))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
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
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v71
						F_errmsg_internal(m, int32(_a_F_ExecAsyncRequest_0), v7+int32(16))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecAsyncRequest_1), int32(44), int32(_a_F_ExecAsyncRequest_3))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
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
		v14 = v9
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
		if v15 != 0 {
			F_InstrStart(m, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v19 = v18
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				if v20 == int32(424) {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+128))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+176))
					m.T0[v25].(func(*base.Module, int32))(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						if v29 == int32(403) {
							F_ExecAsyncAppendResponse(m, l0)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
								if v35 != 0 {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v36 != 0 {
										v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
										if v39&int32(2) != 0 {
											v42 = float64(0)
										} else {
											v42 = float64(1)
										}
										v44 = v42
									} else {
										v44 = float64(0)
									}
									F_InstrStopNode(m, v35, v44)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										m.G0 = v7 + int32(32)
										return
									}
								} else {
									m.G0 = v7 + int32(32)
									return
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v56
								F_errmsg_internal(m, int32(_a_F_ExecAsyncRequest_0), v7)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ExecAsyncRequest_1), int32(128), int32(_a_F_ExecAsyncRequest_2))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
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
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v71
						F_errmsg_internal(m, int32(_a_F_ExecAsyncRequest_0), v7+int32(16))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecAsyncRequest_1), int32(44), int32(_a_F_ExecAsyncRequest_3))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
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
		} else {
			v19 = v14
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == int32(424) {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+128))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+176))
				m.T0[v25].(func(*base.Module, int32))(m, l0)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
					if v29 == int32(403) {
						F_ExecAsyncAppendResponse(m, l0)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
							if v35 != 0 {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v36 != 0 {
									v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
									if v39&int32(2) != 0 {
										v42 = float64(0)
									} else {
										v42 = float64(1)
									}
									v44 = v42
								} else {
									v44 = float64(0)
								}
								F_InstrStopNode(m, v35, v44)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									m.G0 = v7 + int32(32)
									return
								}
							} else {
								m.G0 = v7 + int32(32)
								return
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v56
							F_errmsg_internal(m, int32(_a_F_ExecAsyncRequest_0), v7)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ExecAsyncRequest_1), int32(128), int32(_a_F_ExecAsyncRequest_2))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
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
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v71
					F_errmsg_internal(m, int32(_a_F_ExecAsyncRequest_0), v7+int32(16))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ExecAsyncRequest_1), int32(44), int32(_a_F_ExecAsyncRequest_3))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
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
}
func F_asyncQueueErrdetailForIoError(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v8
	v13 = F_errdetail(m, int32(_a_F_asyncQueueErrdetailForIoError_0), v6)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v13
	}
}
func F_asyncQueueReadAllNotifications(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v193 int64
	_ = v193
	var v194 int64
	_ = v194
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v228 int64
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int64
	_ = v255
	var v256 int64
	_ = v256
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v321 int64
	_ = v321
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int64
	_ = v344
	var v346 int64
	_ = v346
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	v17 = m.G0
	v19 = v17 - int32(_a_F_asyncQueueReadAllNotifications_0)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[0]))
	v26 = F_LWLockAcquire(m, v22+int32(3456), int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v28 = int32(_a_F_asyncQueueReadAllNotifications_1)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[1]))
	v30 = int32(_a_F_asyncQueueReadAllNotifications_2)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[2]))
	v32 = int32(40)
	v35 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29+v31*v32)+96)) = uint8(v35)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[1]))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[2]))
	v43 = v38 + v40*v32
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v44
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v43)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
	if base.B2i32(v48 != base.I32_wrap_i64(v44))|base.B2i32(v46 != v51) == v35 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v19 + int32(_a_F_asyncQueueReadAllNotifications_0)
	return
L4:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[0]))
	F_LWLockRelease(m, v57+int32(3456))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v62 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+97)) = uint8(v62)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[0]))
	F_LWLockRelease(m, v65+int32(3456))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L3
L8:
	;
	v70 = F_GetLatestSnapshot(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v72 = F_RegisterSnapshot(m, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v74 = int32(_a_F_asyncQueueReadAllNotifications_3)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[3])))
	v77 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[3])) = uint8(v77)
	v93 = v46
	goto L11
L11:
	;
	v96 = F_SimpleLruReadPage_ReadOnly(m, int32(_a_F_asyncQueueReadAllNotifications_4), v93, v19)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[0]))
	v335 = F_LWLockAcquire(m, v331+int32(3456), int32(1))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L57
	}
L13:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[4]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v96<<(uint(int32(2))%32))))
	v110 = v19 + int32(16)
	goto L14
L14:
	;
	v123 = int32(0)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	if base.B2i32(v48 == v125)&base.B2i32(v51 == v127) != 0 {
		v246 = v110
		v247 = v123
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[4]))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+28))
	v255 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[5])))
	v256 = base.I64_rem_s(v93, v255)
	F_LWLockRelease(m, v253+base.I32_wrap_i64(v256)<<(uint(int32(7))%32))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L39
	}
L16:
	;
	goto L15
L17:
	;
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v132 = v125 + v104
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v134 = v131 + v133
	v136 = v134 - int32(_a_F_asyncQueueReadAllNotifications_5)
	v138 = base.B2i32(base.Ui32(v136) < base.Ui32(int32(-8193)))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v130 + base.I64_extend_i32_u(v138)
	if base.Ui32(v136) < base.Ui32(int32(-8193)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v143 = int32(0)
	goto L20
L19:
	;
	v143 = v134
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v143
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[6]))
	if v145 != v147 {
		v241 = v110
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if base.Ui32(int32(-8194)) < base.Ui32(v136) {
		v110 = v241
		goto L14
	} else {
		goto L38
	}
L22:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	v150 = F_XidInMVCCSnapshot(m, v149, v72)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v150 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v125
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v127
	v246 = v110
	v247 = int32(1)
	goto L16
L25:
	;
	goto L26
L26:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[7]))
	if v157 == int32(0) {
		v241 = v110
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v161)+8))
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v161)+808))
	if v163 != int64(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v228 == int64(0) {
		v241 = v110
		goto L21
	} else {
		goto L32
	}
L29:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v161)+752))
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v161)+728))
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v161)+704))
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v161)+680))
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v161)+656))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v161)+632))
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v161)+608))
	v173 = *(*int64)(unsafe.Add(mBase, uint32(v161)+584))
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v161)+560))
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v161)+536))
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v161)+512))
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v161)+488))
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v161)+464))
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v161)+440))
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v161)+416))
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v161)+392))
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v161)+368))
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v161)+344))
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v161)+320))
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v161)+296))
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v161)+272))
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v161)+248))
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v161)+224))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v161)+200))
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v161)+176))
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v161)+152))
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v161)+128))
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v161)+104))
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v161)+80))
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v161)+56))
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v161)+32))
	v228 = v166 + (v167 + (v168 + (v169 + (v170 + (v171 + (v172 + (v173 + (v174 + (v175 + (v176 + (v177 + (v178 + (v179 + (v180 + (v181 + (v182 + (v183 + (v184 + (v185 + (v186 + (v187 + (v188 + (v189 + (v190 + (v191 + (v192 + (v193 + (v194 + (v195 + (v196 + v162))))))))))))))))))))))))))))))
	goto L31
L30:
	;
	v228 = v162
	goto L31
L31:
	;
	goto L28
L32:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v132)+8))
	v232 = F_TransactionIdDidCommit(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v232 == int32(0) {
		v241 = v110
		goto L21
	} else {
		goto L34
	}
L34:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v236 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	base.MemoryCopy(m, v110, v132, v236)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v241 = v110 + v238
	goto L21
L38:
	;
	v246 = v241
	v247 = v123
	goto L16
L39:
	;
	v264 = v19 + int32(16)
	if base.Ui32(v264) < base.Ui32(v246) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v267 = v264
	goto L43
L41:
	;
	goto L42
L42:
	;
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	if v51 == v321 {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[7]))
	if v283 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L42
L45:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v303 = v267 + v302
	if base.Ui32(v303) < base.Ui32(v246) {
		v267 = v303
		goto L43
	} else {
		goto L50
	}
L46:
	;
	v287 = v267 + int32(16)
	v288 = int32(0)
	v290 = F_hash_search(m, v283, v287, v288, v288)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	if v290 == int32(0) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v294 = F_strlen(m, v287)
	mBase = m.M
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v267)+12))
	F_NotifyMyFrontEnd(m, v287, v294+v287+int32(1), v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L45
L50:
	;
	goto L44
L51:
	;
	goto L12
L52:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v247|base.B2i32(v323 == v48) == int32(0) {
		v93 = v321
		goto L11
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v247 == int32(0) {
		v93 = v321
		goto L11
	} else {
		goto L56
	}
L55:
	;
	goto L51
L56:
	;
	goto L51
L57:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[1]))
	v339 = int32(_a_F_asyncQueueReadAllNotifications_2)
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[2]))
	v341 = int32(40)
	v343 = v338 + v340*v341
	v344 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v343)+88)) = v344
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(v343)+80)) = v346
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[2]))
	v353 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v338+v349*v341)+97)) = uint8(v353)
	v356 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[0]))
	F_LWLockRelease(m, v356+int32(3456))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_asyncQueueReadAllNotifications[3])) = uint8(v75)
	F_UnregisterSnapshot(m, v72)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	goto L3
}
