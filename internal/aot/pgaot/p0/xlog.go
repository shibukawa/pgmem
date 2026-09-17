package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RemoveXlogFile(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int64
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v8 = m.G0
	v10 = v8 - int32(1072)
	m.G0 = v10
	v13 = l0 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v13
	v16 = v10 + int32(48)
	v21 = F_pg_snprintf(m, v16, int32(1024), int32(_a_F_RemoveXlogFile_0), v10+int32(32))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RemoveXlogFile[0])))
		if v24 != int32(1) {
			v72 = F_errstart(m, int32(13), int32(0))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return
			} else {
				if v72 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
					F_errmsg_internal(m, int32(_a_F_RemoveXlogFile_1), v10)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_RemoveXlogFile_2), int32(4045), int32(_a_F_RemoveXlogFile_3))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							v86 = F_durable_unlink(m, v10+int32(48), int32(15))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								if v86 != 0 {
									m.G0 = v10 + int32(1072)
									return
								} else {
									v88 = int32(_a_F_RemoveXlogFile_4)
									v90 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1]))
									*(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1])) = v90 + int32(1)
									F_XLogArchiveCleanup(m, v13)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										m.G0 = v10 + int32(1072)
										return
									}
								}
							}
						}
					}
				} else {
					v86 = F_durable_unlink(m, v10+int32(48), int32(15))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return
					} else {
						if v86 != 0 {
							m.G0 = v10 + int32(1072)
							return
						} else {
							v88 = int32(_a_F_RemoveXlogFile_4)
							v90 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1]))
							*(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1])) = v90 + int32(1)
							F_XLogArchiveCleanup(m, v13)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								m.G0 = v10 + int32(1072)
								return
							}
						}
					}
				}
			}
		} else {
			v27 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
			if base.Ui64(l1) < base.Ui64(v27) {
				v72 = F_errstart(m, int32(13), int32(0))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					if v72 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
						F_errmsg_internal(m, int32(_a_F_RemoveXlogFile_1), v10)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_RemoveXlogFile_2), int32(4045), int32(_a_F_RemoveXlogFile_3))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								v86 = F_durable_unlink(m, v10+int32(48), int32(15))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									if v86 != 0 {
										m.G0 = v10 + int32(1072)
										return
									} else {
										v88 = int32(_a_F_RemoveXlogFile_4)
										v90 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1]))
										*(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1])) = v90 + int32(1)
										F_XLogArchiveCleanup(m, v13)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											m.G0 = v10 + int32(1072)
											return
										}
									}
								}
							}
						}
					} else {
						v86 = F_durable_unlink(m, v10+int32(48), int32(15))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							if v86 != 0 {
								m.G0 = v10 + int32(1072)
								return
							} else {
								v88 = int32(_a_F_RemoveXlogFile_4)
								v90 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1]))
								*(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1])) = v90 + int32(1)
								F_XLogArchiveCleanup(m, v13)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									m.G0 = v10 + int32(1072)
									return
								}
							}
						}
					}
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[2]))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+320)))
				if v31 != int32(1) {
					v72 = F_errstart(m, int32(13), int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						if v72 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
							F_errmsg_internal(m, int32(_a_F_RemoveXlogFile_1), v10)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_RemoveXlogFile_2), int32(4045), int32(_a_F_RemoveXlogFile_3))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return
								} else {
									v86 = F_durable_unlink(m, v10+int32(48), int32(15))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										if v86 != 0 {
											m.G0 = v10 + int32(1072)
											return
										} else {
											v88 = int32(_a_F_RemoveXlogFile_4)
											v90 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1]))
											*(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1])) = v90 + int32(1)
											F_XLogArchiveCleanup(m, v13)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												m.G0 = v10 + int32(1072)
												return
											}
										}
									}
								}
							}
						} else {
							v86 = F_durable_unlink(m, v10+int32(48), int32(15))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								if v86 != 0 {
									m.G0 = v10 + int32(1072)
									return
								} else {
									v88 = int32(_a_F_RemoveXlogFile_4)
									v90 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1]))
									*(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1])) = v90 + int32(1)
									F_XLogArchiveCleanup(m, v13)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										m.G0 = v10 + int32(1072)
										return
									}
								}
							}
						}
					}
				} else {
					v36 = F_get_dirent_type(m, v16, l0, int32(0), int32(13))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						if v36 != int32(2) {
							v72 = F_errstart(m, int32(13), int32(0))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return
							} else {
								if v72 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
									F_errmsg_internal(m, int32(_a_F_RemoveXlogFile_1), v10)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_RemoveXlogFile_2), int32(4045), int32(_a_F_RemoveXlogFile_3))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											v86 = F_durable_unlink(m, v10+int32(48), int32(15))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												if v86 != 0 {
													m.G0 = v10 + int32(1072)
													return
												} else {
													v88 = int32(_a_F_RemoveXlogFile_4)
													v90 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1]))
													*(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1])) = v90 + int32(1)
													F_XLogArchiveCleanup(m, v13)
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														m.G0 = v10 + int32(1072)
														return
													}
												}
											}
										}
									}
								} else {
									v86 = F_durable_unlink(m, v10+int32(48), int32(15))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										if v86 != 0 {
											m.G0 = v10 + int32(1072)
											return
										} else {
											v88 = int32(_a_F_RemoveXlogFile_4)
											v90 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1]))
											*(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1])) = v90 + int32(1)
											F_XLogArchiveCleanup(m, v13)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												m.G0 = v10 + int32(1072)
												return
											}
										}
									}
								}
							}
						} else {
							v41 = F_InstallXLogFileSegment(m, l2, v16, int32(1), l1, l3)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								if v41 == int32(0) {
									v72 = F_errstart(m, int32(13), int32(0))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										if v72 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
											F_errmsg_internal(m, int32(_a_F_RemoveXlogFile_1), v10)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_RemoveXlogFile_2), int32(4045), int32(_a_F_RemoveXlogFile_3))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return
												} else {
													v86 = F_durable_unlink(m, v10+int32(48), int32(15))
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return
													} else {
														if v86 != 0 {
															m.G0 = v10 + int32(1072)
															return
														} else {
															v88 = int32(_a_F_RemoveXlogFile_4)
															v90 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1]))
															*(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1])) = v90 + int32(1)
															F_XLogArchiveCleanup(m, v13)
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
																return
															} else {
																m.G0 = v10 + int32(1072)
																return
															}
														}
													}
												}
											}
										} else {
											v86 = F_durable_unlink(m, v10+int32(48), int32(15))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												if v86 != 0 {
													m.G0 = v10 + int32(1072)
													return
												} else {
													v88 = int32(_a_F_RemoveXlogFile_4)
													v90 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1]))
													*(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[1])) = v90 + int32(1)
													F_XLogArchiveCleanup(m, v13)
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														m.G0 = v10 + int32(1072)
														return
													}
												}
											}
										}
									}
								} else {
									v47 = F_errstart(m, int32(13), int32(0))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										if v47 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v13
											F_errmsg_internal(m, int32(_a_F_RemoveXlogFile_5), v10+int32(16))
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_RemoveXlogFile_2), int32(4033), int32(_a_F_RemoveXlogFile_3))
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return
												} else {
													v60 = int32(_a_F_RemoveXlogFile_6)
													v62 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[3]))
													*(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[3])) = v62 + int32(1)
													v66 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
													*(*int64)(unsafe.Add(mBase, uint32(l2))) = v66 + int64(1)
													F_XLogArchiveCleanup(m, v13)
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														m.G0 = v10 + int32(1072)
														return
													}
												}
											}
										} else {
											v60 = int32(_a_F_RemoveXlogFile_6)
											v62 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[3]))
											*(*int32)(unsafe.Add(mBase, _c_F_RemoveXlogFile[3])) = v62 + int32(1)
											v66 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
											*(*int64)(unsafe.Add(mBase, uint32(l2))) = v66 + int64(1)
											F_XLogArchiveCleanup(m, v13)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												m.G0 = v10 + int32(1072)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
