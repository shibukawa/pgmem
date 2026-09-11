package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tbm_create_pagetable(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = F_MemoryContextAllocZero(m, v10, int32(32))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v10
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
		if v16 == int32(0) {
			v21 = F_MemoryContextAllocExtended(m, v10, int32(_a_F_tbm_create_pagetable_0), int32(5))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v36 = v21
				*(*int64)(unsafe.Add(mBase, uint32(v12)+12)) = int64(987842478335)
				*(*int64)(unsafe.Add(mBase, uint32(v12))) = int64(256)
				*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v36
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v43 == int32(1) {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v49 = F_pagetable_insert(m, v12, v46, v8+int32(15))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)))
						v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v49))) = v52
						v54 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v54
						v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v56
						v60 = *(*int64)(unsafe.Add(mBase, uint32(l0-int32(-64))))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = v60
						v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+32)) = v62
						v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+40)) = v64
						*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)) = uint8(v51)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
						m.G0 = v8 + int32(16)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
					m.G0 = v8 + int32(16)
					return
				}
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v23
			v27 = F_dsa_allocate_extended(m, v16, int32(_a_F_tbm_create_pagetable_1), int32(5))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v27
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
				v31 = F_dsa_get_address(m, v30, v27)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v36 = v31 + int32(4)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+12)) = int64(987842478335)
					*(*int64)(unsafe.Add(mBase, uint32(v12))) = int64(256)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v36
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v43 == int32(1) {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v49 = F_pagetable_insert(m, v12, v46, v8+int32(15))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)))
							v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v49))) = v52
							v54 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v54
							v56 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v56
							v60 = *(*int64)(unsafe.Add(mBase, uint32(l0-int32(-64))))
							*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = v60
							v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
							*(*int64)(unsafe.Add(mBase, uint32(v49)+32)) = v62
							v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
							*(*int64)(unsafe.Add(mBase, uint32(v49)+40)) = v64
							*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)) = uint8(v51)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
							m.G0 = v8 + int32(16)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_tbm_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+112))
		if v8 == int32(0) {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
			F_pfree(m, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				F_pfree(m, v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					if v27 != 0 {
						F_pfree(m, v27)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							if v30 != 0 {
								F_pfree(m, v30)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						if v30 != 0 {
							F_pfree(m, v30)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							F_pfree(m, l0)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)+100))
			if v14 == int32(0) {
				F_pfree(m, v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					if v27 != 0 {
						F_pfree(m, v27)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							if v30 != 0 {
								F_pfree(m, v30)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						if v30 != 0 {
							F_pfree(m, v30)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							F_pfree(m, l0)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				F_dsa_free(m, v8, v14)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = int32(0)
					F_pfree(m, v6)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						if v27 != 0 {
							F_pfree(m, v27)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								if v30 != 0 {
									F_pfree(m, v30)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v34 = m.ExcPending
										if v34 != 0 {
											return
										} else {
											return
										}
									}
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							if v30 != 0 {
								F_pfree(m, v30)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
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
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		if v27 != 0 {
			F_pfree(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				if v30 != 0 {
					F_pfree(m, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
			if v30 != 0 {
				F_pfree(m, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_tbm_private_iterate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v11 <= v9 {
		v70 = v11
		v73 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v70 <= v73 {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = v13
	v19 = v9
	goto L3
L3:
	;
	v22 = int32(256)
	if v18 <= v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v70 = v66
	v73 = v64
	goto L1
L5:
	;
	v25 = v22
	goto L7
L6:
	;
	v25 = v18
	goto L7
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v19<<(uint(int32(2))%32))))
	v37 = v18
	goto L9
L8:
	;
	v60 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v60
	v64 = v19 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v64 < v66 {
		v18 = v60
		v19 = v64
		goto L3
	} else {
		goto L14
	}
L9:
	;
	if v37 == v25 {
		goto L8
	} else {
		goto L11
	}
L10:
	;
	if int32(255) < v37 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	v42 = int32(1)
	v45 = base.I32_div_s(v37, int32(32))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(8)+v45<<(uint(int32(2))%32))))
	if int32(base.Ui32(v49)>>(uint(v37)%32))&v42 == int32(0) {
		v37 = v37 + v42
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v37
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v70 = v58
	v73 = v19
	goto L1
L14:
	;
	goto L4
L15:
	;
	if v108 < v109 {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v108 = v78
	v109 = v77
	goto L15
L17:
	;
	goto L18
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v73<<(uint(int32(2))%32))))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = v79 + v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v87 < v88 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90+v87<<(uint(int32(2))%32))))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if base.Ui32(v95) <= base.Ui32(v86) {
		v108 = v87
		v109 = v88
		goto L15
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	v99 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v99)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v86
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v103 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v102 + v103
	return v103
L22:
	;
	goto L21
L23:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v112 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(-1)
	return int32(0)
L26:
	;
	v122 = v10 + int32(40)
	goto L28
L27:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117+v108<<(uint(int32(2))%32))))
	v122 = v121
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v125 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v125)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v124
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v128)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v131 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v130 + v131
	return v131
}
func F_tbm_shared_comparator(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = int32(48)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2+v4*v5)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2+v9*v5)))
	return base.B2i32(base.Ui32(v13) < base.Ui32(v8)) - base.B2i32(base.Ui32(v8) < base.Ui32(v13))
}
